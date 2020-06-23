package config

import (
	"fmt"
	"net/url"
)

var (
	// ExchangeURL -
	ExchangeURL = "https://exchange.com"

	// AppName -
	AppName = "exampleapp"

	// Scopes -
	Scopes = "profile"

	// URLRedirect -
	URLRedirect = mountURL()

	// Redirect -
	Redirect = "https://example.com"

	// AppSecret -
	AppSecret = "CLIENT CREDENTIALS"
)

func mountURL() string {
	params := make(url.Values)
	params.Add("client_id", AppName)
	params.Add("redirect_uri", Redirect)
	params.Add("response_type", "code")
	params.Add("scope", Scopes)
	params.Add("state", "hulaqtskuualfbxmrnhbmvji")
	return fmt.Sprintf("%s/api/v3/oauth/oauth2/auth?%s", ExchangeURL, params.Encode())
}
