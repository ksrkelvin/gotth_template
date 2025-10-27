package service

import (
	"encontradev/internal/dto"
	"encontradev/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (s *Service) UpdateUser(ctx *gin.Context, userUpdate dto.UserUpdateRequest) (userDto dto.UserResponse, err error) {
	userSection := s.GetUserFromContext(ctx)

	if userUpdate.Name != "" {
		userSection.Name = userUpdate.Name

	}
	if userUpdate.Avatar != "" {
		userSection.Avatar = userUpdate.Avatar

	}

	user := models.User{
		Model: gorm.Model{
			ID:        userSection.ID,
			UpdatedAt: time.Now(),
		},
		Email:  userSection.Email,
		Name:   userSection.Name,
		Avatar: userSection.Avatar,
	}

	err = s.Repository.UpdateUser(user)
	if err != nil {
		return userDto, err
	}

	return userSection, nil
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
