package playlist

import (
	"encoding/json"
	"errors"
	"github.com/google/go-querystring/query"
	"github.com/najemi-software/tidal-dl/v4/generics"
	"github.com/najemi-software/tidal-dl/v4/requests"
	"github.com/najemi-software/tidal-dl/v4/track"
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

type GetTracksParams struct {
	Offset *int `url:"offset,omitempty"`
	Limit  *int `url:"limit,omitempty"`
}

func GetTracks(id string, params *GetTracksParams) (*TracksResponse, error) {
	paramsValues, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	response, responseData, err := requests.SendBasicRequest(requests.GET, "playlists/"+id+"/tracks", paramsValues, nil, nil)
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
