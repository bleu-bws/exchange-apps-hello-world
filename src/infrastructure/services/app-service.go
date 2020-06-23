package services

import (
	"app-example/src/config"
	"app-example/src/domain/entities"
	"app-example/src/domain/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type appService struct {
	Code       string
	AcessToken entities.Token
}

// NewAppService -
func NewAppService(code string) services.AppService {
	return &appService{
		Code: code,
	}
}

// GetAcessToken -
func (a *appService) GetAcessToken() (err error) {

	params := make(url.Values)
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", config.AppName)
	params.Add("scope", config.Scopes)
	params.Add("client_secret", config.AppSecret)
	params.Add("code", a.Code)
	params.Add("redirect_uri", config.Redirect)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v3/oauth/oauth2/token", config.ExchangeURL), strings.NewReader(params.Encode()))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &a.AcessToken)
	if err != nil {
		return
	}

	return
}

// AccessEndpoint -
func (a *appService) AccessEndpoint(url string) (response []byte, err error) {

	err = a.GetAcessToken()
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.AcessToken.AccessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
