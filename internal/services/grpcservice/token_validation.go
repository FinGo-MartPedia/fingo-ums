package grpcservice

import (
	"context"

	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/pkg/errors"
)

type TokenValidationService struct {
	UserRepository interfaces.IUserRepository
}

func NewTokenValidationService(repo interfaces.IUserRepository) *TokenValidationService {
	return &TokenValidationService{UserRepository: repo}
}

func (s *TokenValidationService) TokenValidation(ctx context.Context, accessToken string) (*helpers.ClaimToken, error) {
	var (
		claimToken *helpers.ClaimToken
		err        error
	)

	claimToken, err = helpers.ValidateToken(ctx, accessToken)
	if err != nil {
		return claimToken, errors.Wrap(err, "failed to validate token")
	}

	_, err = s.UserRepository.GetUserSessionByAccessToken(ctx, accessToken)
	if err != nil {
		return claimToken, errors.Wrap(err, "failed to get user sesion")
	}

	return claimToken, nil
}
