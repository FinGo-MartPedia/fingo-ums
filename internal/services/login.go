package services

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/fingo-martPedia/fingo-ums/internal/models"
	"github.com/fingo-martPedia/fingo-ums/internal/models/requests"
	"github.com/fingo-martPedia/fingo-ums/internal/models/responses"
)

type LoginService struct {
	UserRepository interfaces.IUserRepository
}

func NewLoginService(repo interfaces.IUserRepository) *LoginService {
	return &LoginService{UserRepository: repo}
}

func (s *LoginService) Login(ctx context.Context, request requests.LoginRequest) (responses.LoginResponse, error) {
	var (
		response responses.LoginResponse
		now      = time.Now()
	)

	userDetail, err := s.UserRepository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return response, errors.Wrap(helpers.ErrInvalidCredentials, "failed to get user by username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password)); err != nil {
		return response, errors.Wrap(helpers.ErrInvalidCredentials, "failed to compare password")
	}

	accessToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.Fullname, "access_token", userDetail.Email, now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.Fullname, "refresh_token", userDetail.Email, now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate refresh token")
	}

	userSession := &models.UserSession{
		UserID:              userDetail.ID,
		UserAgent:           request.UserAgent,
		IPAddress:           request.IPAddress,
		AccessToken:         accessToken,
		RefreshToken:        refreshToken,
		AccessTokenExpired:  now.Add(helpers.MapTypeToken["access_token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}
	err = s.UserRepository.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return response, errors.Wrap(err, "failed to insert new session")
	}

	response.AccessToken = accessToken
	response.RefreshToken = refreshToken

	return response, nil
}
