package auth

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v2/requests"
	"net/url"
	"strconv"
)

type LoginLink struct {
	DeviceCode              string `json:"deviceCode"`
	UserCode                string `json:"userCode"`
	VerificationUri         string `json:"verificationUri"`
	VerificationUriComplete string `json:"verificationUriComplete"`
	ExpiresIn               int    `json:"expiresIn"`
	Interval                int    `json:"interval"`
}

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

type User struct {
	UserId             int     `json:"userId"`
	Email              string  `json:"email"`
	CountryCode        string  `json:"countryCode"`
	FullName           *string `json:"fullName"`
	FirstName          *string `json:"firstName"`
	LastName           *string `json:"lastName"`
	NickName           *string `json:"nickName"`
	Username           string  `json:"username"`
	Address            *string `json:"address"`
	PostalCode         *string `json:"postalcode"`
	UseState           *string `json:"useState"`
	PhoneNumber        *string `json:"phoneNumber"`
	Birthday           *string `json:"birthday"`
	ChannelId          int     `json:"channelId"`
	ParentId           int     `json:"parentId"`
	AcceptedEULA       bool    `json:"acceptedEULA"`
	Created            int     `json:"created"`
	Updated            int     `json:"updated"`
	FacebookUId        *int    `json:"facebookUId"`
	AppleUId           *int    `json:"appleUId"`
	GoogleUId          *int    `json:"googleUId"`
	AccountLinkCreated bool    `json:"accountLinkCreated"`
	EmailVerified      bool    `json:"emailVerified"`
	NewUser            bool    `json:"newUser"`
}

type LoginLinkError struct {
	Status           int    `json:"status"`
	ErrorType        string `json:"error"`
	SubStatus        int    `json:"sub_status"`
	ErrorDescription string `json:"error_description"`
}

func (e *LoginLinkError) Error() string {
	return e.ErrorDescription
}

type GrantResponse struct {
	Scope        string `json:"scope"`
	User         User   `json:"user"`
	ClientName   string `json:"clientName"`
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	UserId       int    `json:"user_id"`
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
