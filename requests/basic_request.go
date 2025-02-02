package requests

import (
	"github.com/najemi-software/tidal-dl/credentials"
	"net/http"
	"net/url"
)

func SendBasicRequest(method Method, path string, params url.Values, data url.Values, headers map[string]string) (*http.Response, *string, error) {
	if params == nil {
		params = url.Values{}
	}

	if credentials.SessionId != "" {
		params.Set("sessionId", credentials.SessionId)
	}
	if credentials.CountryCode != "" {
		params.Set("countryCode", credentials.CountryCode)
	}

	if headers == nil {
		headers = map[string]string{}
	}

	if credentials.TokenType != "" && credentials.AccessToken != "" {
		headers["Authorization"] = credentials.TokenType + " " + credentials.AccessToken
	}

	return SendRequest(method, "https://api.tidal.com/v1/"+path, params, data, headers)
}
