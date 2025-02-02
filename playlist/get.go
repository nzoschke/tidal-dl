package playlist

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/requests"
	"strconv"
)

func Get(id string) (*Playlist, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "playlists/"+id, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get playlist. Status: " + strconv.Itoa(response.StatusCode))
	}

	var playlist Playlist

	err = json.Unmarshal([]byte(*responseData), &playlist)
	if err != nil {
		return nil, err
	}

	return &playlist, nil
}
