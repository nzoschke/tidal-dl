package track

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/requests"
	"strconv"
)

func GetStreamUrl(id int) (*StreamUrl, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "tracks/"+strconv.Itoa(id)+"/streamUrl", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get track stream url. Status: " + strconv.Itoa(response.StatusCode))
	}

	var streamUrl StreamUrl

	err = json.Unmarshal([]byte(*responseData), &streamUrl)
	if err != nil {
		return nil, err
	}

	return &streamUrl, nil
}
