package album

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v3/requests"
	"strconv"
)

func Get(id int) (*Album, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "albums/"+strconv.Itoa(id), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get album. Status: " + strconv.Itoa(response.StatusCode))
	}

	var album Album

	err = json.Unmarshal([]byte(*responseData), &album)
	if err != nil {
		return nil, err
	}

	return &album, nil
}
