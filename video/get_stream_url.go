package video

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v2/requests"
	"strconv"
)

func GetStreamUrl(id int) (*StreamUrl, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "videos/"+strconv.Itoa(id)+"/streamUrl", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get video stream url. Status: " + strconv.Itoa(response.StatusCode))
	}

	var streamUrl StreamUrl

	err = json.Unmarshal([]byte(*responseData), &streamUrl)
	if err != nil {
		return nil, err
	}

	return &streamUrl, nil
}
