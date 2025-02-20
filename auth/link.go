package auth

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v3/requests"
	"net/url"
	"strconv"
)

func GetLoginLink() (*LoginLink, error) {
	response, responseData, err := requests.SendRequest(requests.POST, "https://auth.tidal.com/v1/oauth2/device_authorization", nil, url.Values{
		"client_id": {requests.ClientId},
		"scope":     {"r_usr w_usr w_sub"},
	}, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get login link. Status: " + strconv.Itoa(response.StatusCode))
	}

	var loginLink LoginLink

	err = json.Unmarshal([]byte(*responseData), &loginLink)
	if err != nil {
		return nil, err
	}

	return &loginLink, nil
}

func GetLoginLinkStatus(deviceCode string) (*GrantResponse, error) {
	response, responseData, err := requests.SendRequest(requests.POST, "https://auth.tidal.com/v1/oauth2/token", nil, url.Values{
		"client_id":     {requests.ClientId},
		"client_secret": {requests.ClientSecret},
		"device_code":   {deviceCode},
		"grant_type":    {"urn:ietf:params:oauth:grant-type:device_code"},
		"scope":         {"r_usr w_usr w_sub"},
	}, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		var loginLinkStatus GrantResponse

		err = json.Unmarshal([]byte(*responseData), &loginLinkStatus)
		if err != nil {
			return nil, err
		}

		return &loginLinkStatus, nil
	} else if response.StatusCode == 400 {
		var loginLinkError LoginLinkError

		err = json.Unmarshal([]byte(*responseData), &loginLinkError)
		if err != nil {
			return nil, err
		}

		return nil, &loginLinkError
	} else {
		return nil, errors.New("Failed to get login link status. Status: " + strconv.Itoa(response.StatusCode))
	}
}
