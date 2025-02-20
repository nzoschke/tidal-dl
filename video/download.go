package video

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/najemi-software/tidal-dl/v3/requests"
	"github.com/najemi-software/tidal-dl/v3/video_quality"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type VideoStreamUrl struct {
	Codec       string
	M3U8Url     string
	Resolution  string
	Resolutions []int
}

func getSub(s, start, end string) string {
	startIndex := strings.Index(s, start)
	if startIndex == -1 {
		return ""
	}
	startIndex += len(start)
	endIndex := strings.Index(s[startIndex:], end)
	if endIndex == -1 {
		return s[startIndex:]
	}
	return s[startIndex : startIndex+endIndex]
}

func getSubOnlyStart(s, start string) string {
	startIndex := strings.Index(s, start)
	if startIndex == -1 {
		return ""
	}
	return s[startIndex:]
}

func getResolutionList(url string) (*[]VideoStreamUrl, error) {
	var ret []VideoStreamUrl

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	txt := string(body)
	array := strings.Split(txt, "#")

	for _, item := range array {
		if !strings.Contains(item, "RESOLUTION=") || !strings.Contains(item, "EXT-X-STREAM-INF:") {
			continue
		}

		stream := VideoStreamUrl{
			Codec:      getSub(item, "CODECS=\"", "\""),
			M3U8Url:    strings.TrimSpace(getSubOnlyStart(item, "http")),
			Resolution: strings.TrimSpace(getSub(item, "RESOLUTION=", "http")),
		}

		if strings.Contains(stream.Resolution, ",") {
			stream.Resolution = strings.Split(stream.Resolution, ",")[0]
		}

		resStrings := strings.Split(stream.Resolution, "x")

		xInt, err := strconv.Atoi(resStrings[0])
		if err != nil {
			return nil, err
		}

		yInt, err := strconv.Atoi(resStrings[1])
		if err != nil {
			return nil, err
		}

		stream.Resolutions = []int{xInt, yInt}
		ret = append(ret, stream)
	}

	return &ret, nil
}

// Main function to download the .ts files and convert them to mp4
func downloadAndConvertToMp4(m3u8Content string) (*[]byte, error) {
	// Parse the .ts file URLs from the M3U8 content
	lines := strings.Split(m3u8Content, "\n")
	var tsFiles []string
	for _, line := range lines {
		if strings.HasPrefix(line, "https://") && strings.Contains(strings.Split(line, "?")[0], ".ts") {
			tsFiles = append(tsFiles, line)
		}
	}

	r, w := io.Pipe()
	var output bytes.Buffer

	cmd := exec.Command("ffmpeg", "-f", "mpegts", "-i", "-", "-codec", "copy", "-bsf:a", "aac_adtstoasc", "-movflags", "frag_keyframe+empty_moov", "-f", "mp4", "-")
	cmd.Stdin = r
	cmd.Stdout = &output
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	var streamError error
	// Stream .ts files to ffmpeg
	go func() {
		for _, url := range tsFiles {
			resp, err := http.Get(url)
			if err != nil {
				streamError = err
				return
			}

			// Copy directly to ffmpeg stdin
			if _, err := io.Copy(w, resp.Body); err != nil {
				streamError = err
				return
			}

			err = resp.Body.Close()
			if err != nil {
				streamError = err
				return
			}
		}

		err = w.Close()
		if err != nil {
			streamError = err
			return
		}
	}()

	if streamError != nil {
		return nil, streamError
	}

	data := output.Bytes()

	return &data, nil
}

func Download(id int, videoQuality video_quality.VideoQuality) (*[]byte, error) {
	playbackInfo, err := GetPlaybackInfo(id)
	if err != nil {
		return nil, err
	}

	if !strings.Contains(playbackInfo.ManifestMimeType, "vnd.tidal.emu") {
		return nil, errors.New("Failed to get stream. Mimetype is not supported, received: " + playbackInfo.ManifestMimeType)
	}

	manifestJson, err := base64.StdEncoding.DecodeString(playbackInfo.Manifest)
	if err != nil {
		return nil, err
	}

	var manifest Manifest

	err = json.Unmarshal(manifestJson, &manifest)
	if err != nil {
		return nil, err
	}

	url := manifest.Urls[0]

	resolutionList, err := getResolutionList(url)
	if err != nil {
		return nil, err
	}

	var resolution *VideoStreamUrl

	for _, r := range *resolutionList {
		if int(videoQuality) <= r.Resolutions[1] {
			resolution = &r
		}
	}

	if resolution == nil {
		return nil, errors.New("no stream found with the requested quality. Try again with a lower quality")
	}

	response, m3u8Content, err := requests.SendRequest(requests.GET, resolution.M3U8Url, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get stream. Status code: " + strconv.Itoa(response.StatusCode))
	}

	return downloadAndConvertToMp4(*m3u8Content)
}

func DownloadToFile(id int, videoQuality video_quality.VideoQuality, pathWithoutExtension string) error {
	path := pathWithoutExtension + ".mp4"

	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}

	data, err := Download(id, videoQuality)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	if _, err := file.Write(*data); err != nil {
		return err
	}

	return file.Close()
}
