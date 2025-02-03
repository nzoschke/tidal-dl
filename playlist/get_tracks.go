package playlist

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v2/generics"
	"github.com/najemi-software/tidal-dl/v2/requests"
	"github.com/najemi-software/tidal-dl/v2/track"
	"strconv"
)

type Track struct {
	track.Track
	Description interface{} `json:"description"`
	DateAdded   string      `json:"dateAdded"`
	Index       int         `json:"index"`
	ItemUuid    string      `json:"itemUuid"`
}

type TracksResponse = generics.PaginatedResponse[Track]

func GetTracks(id string) (*TracksResponse, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "playlists/"+id+"/tracks", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get playlist tracks. Status: " + strconv.Itoa(response.StatusCode))
	}

	var tracks TracksResponse

	err = json.Unmarshal([]byte(*responseData), &tracks)
	if err != nil {
		return nil, err
	}

	return &tracks, nil
}
