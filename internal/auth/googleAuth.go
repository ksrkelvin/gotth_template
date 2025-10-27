package auth

import (
	"encoding/json"
	"encontradev/internal/dto"
	"encontradev/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func (auth *Auth) GoogleLogin(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			c.String(500, "Erro inesperado: "+err.Error())
		}
	}()

	url := auth.GoogleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(302, url)
}

func (auth *Auth) GoogleCallback(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			c.String(500, "Erro inesperado: "+err.Error())
		}
	}()

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

	profile := dto.ExternalAuthProfile{}

	if err = json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		c.String(500, "Erro ao decodificar perfil do Google")
		return
	}

	profile.Source = "Gooogle"
	userModel, err := auth.SaveOrGetExternalUser(profile)
	if err != nil {
		c.String(500, "Erro ao tentar salvar usuario Google")

	}

	tokenJWT, err := auth.GenerateJWT(userModel.Email, "diino-app")
	if err != nil {
		c.String(500, "Erro ao gerar token JWT")
		return
	}

	c.SetCookie("jwt", tokenJWT, 3600*24, "/", "", false, true)

	c.Redirect(302, "/")

}

func (auth *Auth) SaveOrGetExternalUser(profile dto.ExternalAuthProfile) (userModels models.User, err error) {
	user, err := auth.Repository.GetUserByEmail(profile.Email)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return auth.Repository.CreateExternalUser(profile)
		}
		return userModels, err
	}
	return user, err

}
