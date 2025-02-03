package artist

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v2/generics"
	"github.com/najemi-software/tidal-dl/v2/requests"
	"github.com/najemi-software/tidal-dl/v2/video"
	"strconv"
)

type VideosResponse = generics.PaginatedResponse[video.Video]

func GetVideos(id int) (*VideosResponse, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "artists/"+strconv.Itoa(id)+"/videos", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get artist videos. Status: " + strconv.Itoa(response.StatusCode))
	}

	var videos VideosResponse

	err = json.Unmarshal([]byte(*responseData), &videos)
	if err != nil {
		return nil, err
	}

	return &videos, nil
}
