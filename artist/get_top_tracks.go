package artist

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v3/generics"
	"github.com/najemi-software/tidal-dl/v3/requests"
	"github.com/najemi-software/tidal-dl/v3/track"
	"strconv"
)

type TopTracksResponse = generics.PaginatedResponse[track.Track]

func GetTopTracks(id int) (*TopTracksResponse, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "artists/"+strconv.Itoa(id)+"/toptracks", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get artist top tracks. Status: " + strconv.Itoa(response.StatusCode))
	}

	var topTracks TopTracksResponse

	err = json.Unmarshal([]byte(*responseData), &topTracks)
	if err != nil {
		return nil, err
	}

	return &topTracks, nil
}
