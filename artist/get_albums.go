package artist

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v3/album"
	"github.com/najemi-software/tidal-dl/v3/generics"
	"github.com/najemi-software/tidal-dl/v3/requests"
	"strconv"
)

type AlbumsResponse = generics.PaginatedResponse[album.Album]

func GetAlbums(id int) (*AlbumsResponse, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "artists/"+strconv.Itoa(id)+"/albums", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get artist albums. Status: " + strconv.Itoa(response.StatusCode))
	}

	var albums AlbumsResponse

	err = json.Unmarshal([]byte(*responseData), &albums)
	if err != nil {
		return nil, err
	}

	return &albums, nil
}
