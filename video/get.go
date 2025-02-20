package video

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v3/requests"
	"strconv"
)

func Get(id int) (*Video, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "videos/"+strconv.Itoa(id), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get video. Status: " + strconv.Itoa(response.StatusCode))
	}

	var video Video

	err = json.Unmarshal([]byte(*responseData), &video)
	if err != nil {
		return nil, err
	}

	return &video, nil
}
