package services

import (
	"context"

	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
)

type LogoutService struct {
	UserRepository interfaces.IUserRepository
}

func NewLogoutService(repo interfaces.IUserRepository) *LogoutService {
	return &LogoutService{UserRepository: repo}
}

func (s *LogoutService) Logout(ctx context.Context, accessToken string) error {
	return s.UserRepository.DeleteUserSession(ctx, accessToken)
}
