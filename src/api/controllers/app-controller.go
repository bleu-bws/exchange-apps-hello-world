package controllers

import (
	"encoding/json"
	"exchange-apps-hello-world/src/config"
	"exchange-apps-hello-world/src/domain/entities"
	"exchange-apps-hello-world/src/infrastructure/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppController -
type AppController struct{}

// Home -
func (ctrl AppController) Home(c *gin.Context) {

	code := c.Query("code")

	if code == "" {
		c.Redirect(http.StatusTemporaryRedirect, config.URLRedirect)
		return
	}

	services := services.NewAppService(code)

	data, err := services.AccessEndpoint(fmt.Sprintf("%s/profile/info", config.ExchangeURL))
	if err != nil {
		c.String(500, "")
		return
	}

	profile := entities.Profile{}

	err = json.Unmarshal(data, &profile)
	if err != nil {
		c.String(500, "")
		return
	}

	c.String(200, fmt.Sprintf("Hello %s", profile.Result.Email))

}
