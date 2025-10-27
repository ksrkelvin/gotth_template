package service

import (
	"encontradev/internal/dto"
	"encontradev/internal/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetUser(ctx *gin.Context) (userDto dto.UserResponse, err error) {
	user := s.GetUserFromContext(ctx)

	userResponse := dto.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Avatar: user.Avatar,
	}

	return userResponse, err
}

func (controller *Service) GetUserFromContext(c *gin.Context) (userModel models.User) {
	user, exists := c.Get("user")
	if !exists || user == nil {
		return models.User{}
	}

	u, ok := user.(models.User)
	if !ok {
		return models.User{}
	}

	return u
}
