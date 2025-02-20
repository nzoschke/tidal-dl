package video

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v4/requests"
	"net/url"
	"strconv"
)

func GetPlaybackInfo(id int) (*PlaybackInfo, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "videos/"+strconv.Itoa(id)+"/playbackinfopostpaywall", url.Values{
		"videoquality":      {"HIGH"},
		"playbackmode":      {"STREAM"},
		"assetpresentation": {"FULL"},
	}, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get video playback info. Status: " + strconv.Itoa(response.StatusCode))
	}

	var playbackInfo PlaybackInfo

	err = json.Unmarshal([]byte(*responseData), &playbackInfo)
	if err != nil {
		return nil, err
	}

	return &playbackInfo, nil
}
