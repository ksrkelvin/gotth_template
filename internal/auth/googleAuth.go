package auth

import (
	"encoding/json"
	"encontradev/internal/dto"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (auth *Auth) GoogleLogin(c *gin.Context) {
	url := auth.GoogleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(302, url)
}

func (auth *Auth) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := auth.GoogleOauthConfig.Exchange(c, code)
	if err != nil {
		c.String(500, "Erro ao autenticar com Google")
		return
	}

	client := auth.GoogleOauthConfig.Client(c, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil || resp.StatusCode != 200 {
		c.String(500, "Erro ao buscar perfil do Google")
		return
	}
	defer resp.Body.Close()

	profile := dto.ExternalAuthProfile

	if err = json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		c.String(500, "Erro ao decodificar perfil do Google")
		return
	}

	tokenJWT, err := auth.GenerateJWT(profile.Email, "diino-app")
	if err != nil {
		c.String(500, "Erro ao gerar token JWT")
		return
	}

	c.SetCookie("jwt", tokenJWT, 3600*24, "/", "", false, true)

	c.Redirect(302, "/home")

}
