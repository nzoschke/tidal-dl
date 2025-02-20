package search

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v3/requests"
	"net/url"
	"strconv"
	"strings"
)

func Get(query string, types []Types, limit int, offset int) (*Response, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "search", url.Values{
		"query":  {query},
		"types":  {strings.Join(typesToStrings(types), ",")},
		"limit":  {strconv.Itoa(limit)},
		"offset": {strconv.Itoa(offset)},
	}, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get search results. Status: " + strconv.Itoa(response.StatusCode))
	}

	var results Response

	err = json.Unmarshal([]byte(*responseData), &results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
