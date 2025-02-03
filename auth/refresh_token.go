package auth

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v2/credentials"
	"github.com/najemi-software/tidal-dl/v2/requests"
	"net/url"
	"strconv"
)

func RefreshToken() (*GrantResponse, error) {
	response, responseData, err := requests.SendRequest(requests.POST, "https://auth.tidal.com/v1/oauth2/token", nil, url.Values{
		"client_id":     {requests.ClientId},
		"client_secret": {requests.ClientSecret},
		"refresh_token": {credentials.RefreshToken},
		"grant_type":    {"refresh_token"},
	}, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to refresh token. Status: " + strconv.Itoa(response.StatusCode))
	}

	var grant GrantResponse

	err = json.Unmarshal([]byte(*responseData), &grant)
	if err != nil {
		return nil, err
	}

	return &grant, nil
}
