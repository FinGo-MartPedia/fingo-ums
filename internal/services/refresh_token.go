package services

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/fingo-martPedia/fingo-ums/internal/models/responses"
)

type RefreshTokenService struct {
	UserRepository interfaces.IUserRepository
}

func NewRefreshTokenService(repo interfaces.IUserRepository) *RefreshTokenService {
	return &RefreshTokenService{UserRepository: repo}
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (responses.RefreshTokenResponse, error) {
	var (
		response responses.RefreshTokenResponse
		now      = time.Now()
	)

	accessToken, err := helpers.GenerateToken(ctx, tokenClaim.UserID, tokenClaim.Username, tokenClaim.Fullname, "access_token", tokenClaim.Email, now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate new access token")
	}

	err = s.UserRepository.UpdateTokenByRefreshToken(ctx, accessToken, refreshToken)
	if err != nil {
		return response, errors.Wrap(err, "failed to update access token")
	}

	response.AccessToken = accessToken
	return response, nil
}
