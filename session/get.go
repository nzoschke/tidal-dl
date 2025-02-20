package session

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v4/requests"
	"strconv"
)

type Client struct {
	Id                       int     `json:"id"`
	Name                     string  `json:"name"`
	AuthorizedForOffline     bool    `json:"authorizedForOffline"`
	AuthorizedForOfflineDate *string `json:"authorizedForOfflineDate"`
}

type Session struct {
	SessionId   string `json:"sessionId"`
	UserId      int    `json:"userId"`
	CountryCode string `json:"countryCode"`
	ChannelId   int    `json:"channelId"`
	PartnerId   int    `json:"partnerId"`
	Client      Client `json:"client"`
}

func Get() (*Session, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "sessions", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get session. Status: " + strconv.Itoa(response.StatusCode))
	}

	var session Session

	err = json.Unmarshal([]byte(*responseData), &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
