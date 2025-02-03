package artist

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/requests"
	"strconv"
)

func GetBio(id int) (*Bio, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "artists/"+strconv.Itoa(id)+"/bio", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get artist bio. Status: " + strconv.Itoa(response.StatusCode))
	}

	var bio Bio

	err = json.Unmarshal([]byte(*responseData), &bio)
	if err != nil {
		return nil, err
	}

	return &bio, nil
}
