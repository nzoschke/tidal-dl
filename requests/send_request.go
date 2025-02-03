package requests

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func SendRequest(method Method, uri string, params url.Values, data url.Values, headers map[string]string) (*http.Response, *string, error) {
	if headers == nil {
		headers = map[string]string{}
	}

	// Prepare request body if Data is provided
	var body io.Reader
	if data != nil {
		body = strings.NewReader(data.Encode())
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	if params != nil {
		uri += "?" + params.Encode()
	}

	// Create new HTTP request
	req, err := http.NewRequest(string(method), uri, body)
	if err != nil {
		return nil, nil, err
	}

	// Add headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, nil, err
	}

	responseBodyString := string(responseBody)

	return resp, &responseBodyString, nil
}
