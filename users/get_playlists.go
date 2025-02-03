package users

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v2/generics"
	"github.com/najemi-software/tidal-dl/v2/playlist"
	"github.com/najemi-software/tidal-dl/v2/requests"
	"strconv"
)

type PlaylistsResponse = generics.PaginatedResponse[playlist.Playlist]

func GetPlaylists(id int) (*PlaylistsResponse, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "users/"+strconv.Itoa(id)+"/playlists", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get user playlists. Status: " + strconv.Itoa(response.StatusCode))
	}

	var playlists PlaylistsResponse

	err = json.Unmarshal([]byte(*responseData), &playlists)
	if err != nil {
		return nil, err
	}

	return &playlists, nil
}
