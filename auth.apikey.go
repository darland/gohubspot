package gohubspot

import (
	"net/http"
)

type APIKeyAuth struct {
	apiKey string
}

const (
	apiKeyParam = "hapikey"
)

// NewAPIKeyAuth create new API KEY Authenticator
func NewAPIKeyAuth(apikey string) APIKeyAuth {
	return APIKeyAuth{apiKey: apikey}
}

// Authenticate set auth
func (auth APIKeyAuth) Authenticate(request *http.Request) error {
	params := request.URL.Query()
	params.Set(apiKeyParam, auth.apiKey)
	request.URL.RawQuery = params.Encode()
	return nil
}
