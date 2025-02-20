package artist

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v3/generics"
	"github.com/najemi-software/tidal-dl/v3/requests"
	"strconv"
)

type SimilarArtistsResponse struct {
	generics.PaginatedResponse[SimilarArtist]
	Source string `json:"source"`
}

func GetSimilarArtists(id int) (*SimilarArtistsResponse, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "artists/"+strconv.Itoa(id)+"/similar", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get artist similar artists. Status: " + strconv.Itoa(response.StatusCode))
	}

	var similarArtists SimilarArtistsResponse

	err = json.Unmarshal([]byte(*responseData), &similarArtists)
	if err != nil {
		return nil, err
	}

	return &similarArtists, nil
}
