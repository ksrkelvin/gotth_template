package service

import (
	"encontradev/internal/dto"

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

func (controller *Service) GetUserFromContext(c *gin.Context) (userModel dto.UserResponse) {
	user, exists := c.Get("user")
	if !exists || user == nil {
		return dto.UserResponse{}
	}

	u, ok := user.(dto.UserResponse)
	if !ok {
		return dto.UserResponse{}
	}

	return u
}
