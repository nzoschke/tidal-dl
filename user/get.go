package user

import (
	"encoding/json"
	"errors"
	"github.com/najemi-software/tidal-dl/v3/requests"
	"strconv"
)

func Get(id int) (*User, error) {
	response, responseData, err := requests.SendBasicRequest(requests.GET, "users/"+strconv.Itoa(id), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Failed to get user. Status: " + strconv.Itoa(response.StatusCode))
	}

	var user User

	err = json.Unmarshal([]byte(*responseData), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
