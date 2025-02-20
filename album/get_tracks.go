package album

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v3/generics"
	"github.com/najemi-software/tidal-dl/v3/requests"
	"github.com/najemi-software/tidal-dl/v3/track"
	"strconv"
)

type TracksResponse = generics.PaginatedResponse[track.Track]

func GetTracks(id int) (*TracksResponse, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "albums/"+strconv.Itoa(id)+"/tracks", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get album tracks. Status: " + strconv.Itoa(response.StatusCode))
	}

	var tracks TracksResponse

	err = json.Unmarshal([]byte(*responseData), &tracks)
	if err != nil {
		return nil, err
	}

	return &tracks, nil
}
