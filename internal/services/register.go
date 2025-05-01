package services

import (
	"context"

	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/fingo-martPedia/fingo-ums/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	UserRepository interfaces.IUserRepository
}

func (s *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	request.Password = string(hashedPassword)

	existUser, err := s.UserRepository.GetUserByUsername(ctx, request.Username)
	if existUser.ID != 0 {
		return nil, helpers.ErrUsernameExists
	}

	err = s.UserRepository.InsertNewUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	response := request
	response.Password = ""
	return response, nil
}
