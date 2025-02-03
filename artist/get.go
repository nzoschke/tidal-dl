package artist

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/requests"
	"strconv"
)

func Get(id int) (*Artist, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "artists/"+strconv.Itoa(id), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get artist. Status: " + strconv.Itoa(response.StatusCode))
	}

	var artist Artist

	err = json.Unmarshal([]byte(*responseData), &artist)
	if err != nil {
		return nil, err
	}

	return &artist, nil
}
