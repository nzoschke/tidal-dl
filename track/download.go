package track

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/najemi-software/tidal-dl/v4/audio_quality"
	"github.com/najemi-software/tidal-dl/v4/encryption_type"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func getExtension(url string) string {
	base := strings.Split(url, "?")[0]

	supported := []string{".flac", ".mp4"}
	for _, ext := range supported {
		if strings.HasSuffix(base, ext) {
			return base
		}
	}

	return ".m4a"
}

func getFileSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return -1
	}

	return fileInfo.Size()
}

func getRemoteFileSize(url string) (*int64, error) {
	response, err := http.Head(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get remote file size. Received status: " + string(response.Status))
	}

	headerValue := response.Header.Get("Content-Length")
	if headerValue == "" {
		return nil, errors.New("failed to get remote file size. No Content-Length header provided")
	}

	size, err := strconv.ParseInt(headerValue, 10, 64)
	if err != nil {
		return nil, err
	}

	return &size, nil
}

func requiresDownload(path, url string) (*bool, *int64, error) {
	fileSize := getFileSize(path)

	remoteFileSize, err := getRemoteFileSize(url)
	if err != nil {
		return nil, nil, err
	}

	t := true
	if fileSize == -1 || *remoteFileSize == -1 {
		return &t, remoteFileSize, nil
	} else if fileSize != *remoteFileSize {
		return &t, remoteFileSize, nil
	}

	f := false
	return &f, remoteFileSize, nil
}

type Chunk struct {
	Index int
	Data  []byte
	Err   error
}

func downloadFile(url string, fileSize, partSize int64) (*[]byte, error) {
	partsCount := int(math.Ceil(float64(fileSize) / float64(partSize)))
	chunks := make(chan Chunk, partsCount)
	var wg sync.WaitGroup

	// Spawn a goroutine for each chunk
	for i := 0; i < partsCount; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			start := int64(index) * partSize
			end := start + partSize - 1
			if end >= fileSize {
				end = fileSize - 1
			}

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				chunks <- Chunk{Index: index, Err: err}
				return
			}
			req.Header.Set("Range", "bytes="+strconv.Itoa(int(start))+"-"+strconv.Itoa(int(end)))

			resp, err := http.DefaultClient.Do(req)
			if err != nil || (resp.StatusCode != http.StatusPartialContent && resp.StatusCode != http.StatusOK) {
				chunks <- Chunk{Index: index, Err: fmt.Errorf("failed to download chunk %d: %v", index, err)}
				return
			}

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				chunks <- Chunk{Index: index, Err: err}
				return
			}

			err = resp.Body.Close()
			if err != nil {
				chunks <- Chunk{Index: index, Err: err}
				return
			}

			chunks <- Chunk{Index: index, Data: data}
		}(i)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(chunks)
	}()

	// Combine chunks in the correct order
	combined := make([][]byte, partsCount)
	for chunk := range chunks {
		if chunk.Err != nil {
			return nil, chunk.Err
		}
		combined[chunk.Index] = chunk.Data
	}

	// Write combined data to buffer
	var buffer bytes.Buffer
	for _, data := range combined {
		if _, err := buffer.Write(data); err != nil {
			return nil, err
		}
	}

	data := buffer.Bytes()
	return &data, nil
}

func convertExtensionToM4a(extension, codec string) string {
	if strings.Contains(codec, "ac4") || strings.Contains(codec, "mha1") {
		return extension
	}

	if extension != ".mp4" {
		return extension
	}

	return ".m4a"
}

func Download(id int, audioQuality audio_quality.AudioQuality, pathWithoutExtension string) (*[]byte, *string, error) {
	playbackInfo, err := GetPlaybackInfo(id, audioQuality)
	if err != nil {
		return nil, nil, err
	}

	if !strings.Contains(playbackInfo.ManifestMimeType, "vnd.tidal.bt") {
		return nil, nil, errors.New("Failed to get stream. Mimetype is not supported, received: " + playbackInfo.ManifestMimeType)
	}

	manifestJson, err := base64.StdEncoding.DecodeString(playbackInfo.Manifest)
	if err != nil {
		return nil, nil, err
	}

	var manifest Manifest

	err = json.Unmarshal(manifestJson, &manifest)
	if err != nil {
		return nil, nil, err
	}

	if manifest.EncryptionType != encryption_type.None {
		return nil, nil, errors.New("Failed to get stream. Encryption type is not supported, received: " + string(manifest.EncryptionType))
	}

	url := manifest.Urls[0]

	contentLength, err := getRemoteFileSize(url)
	if err != nil {
		return nil, nil, err
	}

	data, err := downloadFile(url, *contentLength, 1048576)
	if err != nil {
		return nil, nil, err
	}

	extension := convertExtensionToM4a(getExtension(url), manifest.Codecs)

	return data, &extension, err
}

func DownloadToFile(id int, audioQuality audio_quality.AudioQuality, pathWithoutExtension string) error {
	playbackInfo, err := GetPlaybackInfo(id, audioQuality)
	if err != nil {
		return err
	}

	if !strings.Contains(playbackInfo.ManifestMimeType, "vnd.tidal.bt") {
		return errors.New("Failed to get stream. Mimetype is not supported, received: " + playbackInfo.ManifestMimeType)
	}

	manifestJson, err := base64.StdEncoding.DecodeString(playbackInfo.Manifest)
	if err != nil {
		return err
	}

	var manifest Manifest

	err = json.Unmarshal(manifestJson, &manifest)
	if err != nil {
		return err
	}

	if manifest.EncryptionType != encryption_type.None {
		return errors.New("Failed to get stream. Encryption type is not supported, received: " + string(manifest.EncryptionType))
	}

	url := manifest.Urls[0]

	path := pathWithoutExtension + convertExtensionToM4a(getExtension(url), manifest.Codecs)

	downloadReq, contentLength, err := requiresDownload(path, url)
	if err != nil {
		return err
	}

	if !*downloadReq {
		return nil
	}

	data, err := downloadFile(url, *contentLength, 1048576)
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
