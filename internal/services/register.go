package services

import (
	"context"

	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/fingo-martPedia/fingo-ums/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	RegisterRepository interfaces.IRegisterRepository
}

func (s *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	request.Password = string(hashedPassword)

	err = s.RegisterRepository.InsertNewUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	response := request
	response.Password = ""
	return response, nil
}
