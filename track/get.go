package track

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/requests"
	"strconv"
)

func Get(id string) (*Track, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "tracks/"+id, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get track. Status: " + strconv.Itoa(response.StatusCode))
	}

	var track Track

	err = json.Unmarshal([]byte(*responseData), &track)
	if err != nil {
		return nil, err
	}

	return &track, nil
}
