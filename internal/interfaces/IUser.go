package interfaces

import (
	"context"

	"github.com/fingo-martPedia/fingo-ums/internal/models"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	InsertNewUserSession(ctx context.Context, userSession *models.UserSession) error
	GetUserSessionByAccessToken(ctx context.Context, accessToken string) (models.UserSession, error)
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error)
	DeleteUserSession(ctx context.Context, accessToken string) error
	UpdateTokenByRefreshToken(ctx context.Context, accessToken string, refreshToken string) error
}
