package tidal_dl

import (
	"errors"
	"fmt"
	"github.com/najemi-software/tidal-dl/v2/auth"
	"github.com/najemi-software/tidal-dl/v2/credentials"
	"github.com/najemi-software/tidal-dl/v2/session"
	"time"
)

func InteractiveLogin() error {
	link, err := auth.GetLoginLink()
	if err != nil {
		return err
	}

	fmt.Println("Visit the following link to login: https://" + link.VerificationUriComplete)

	startTime := time.Now()
	for startTime.Add(time.Second * time.Duration(link.ExpiresIn-3)).After(time.Now()) {
		status, err := auth.GetLoginLinkStatus(link.DeviceCode)
		if err != nil {
			var loginLinkError *auth.LoginLinkError
			if !errors.As(err, &loginLinkError) {
				return err
			}

			if loginLinkError.ErrorType == "authorization_pending" {
				time.Sleep(time.Second * time.Duration(link.Interval))
				break
			}

			return err
		}

		credentials.TokenType = status.TokenType
		credentials.AccessToken = status.AccessToken
		credentials.RefreshToken = status.RefreshToken
		credentials.Expires = time.Now().Add(time.Second * time.Duration(status.ExpiresIn))

		loginSession, err := session.Get()
		if err != nil {
			return err
		}

		credentials.SessionId = loginSession.SessionId
		credentials.CountryCode = loginSession.CountryCode

		return nil
	}

	return errors.New("login link expired")
}

func RefreshTokenLogin(refreshToken string) error {
	credentials.RefreshToken = refreshToken

	grant, err := auth.RefreshToken()
	if err != nil {
		return err
	}

	credentials.TokenType = grant.TokenType
	credentials.AccessToken = grant.AccessToken
	credentials.RefreshToken = grant.RefreshToken
	credentials.Expires = time.Now().Add(time.Second * time.Duration(grant.ExpiresIn))

	loginSession, err := session.Get()
	if err != nil {
		return err
	}

	credentials.SessionId = loginSession.SessionId
	credentials.CountryCode = loginSession.CountryCode

	return nil
}

func AccessTokenLogin(tokenType, accessToken string) error {
	credentials.TokenType = tokenType
	credentials.AccessToken = accessToken

	loginSession, err := session.Get()
	if err != nil {
		return err
	}

	credentials.SessionId = loginSession.SessionId
	credentials.CountryCode = loginSession.CountryCode

	return nil
}
