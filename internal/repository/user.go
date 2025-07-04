package repository

import (
	"context"
	"errors"

	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/fingo-martPedia/fingo-ums/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

var _ interfaces.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var (
		user models.User
		err  error
	)
	err = r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, userSession *models.UserSession) error {
	return r.DB.Create(&userSession).Error
}

func (r *UserRepository) GetUserSessionByAccessToken(ctx context.Context, accessToken string) (models.UserSession, error) {
	var userSession models.UserSession
	err := r.DB.Where("access_token = ?", accessToken).First(&userSession).Error
	if err != nil {
		return userSession, err
	}

	if userSession.ID == 0 {
		return userSession, errors.New("user session not found")
	}
	return userSession, nil
}

func (r *UserRepository) GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error) {
	var userSession models.UserSession
	err := r.DB.Where("refresh_token = ?", refreshToken).First(&userSession).Error
	if err != nil {
		return userSession, err
	}

	if userSession.ID == 0 {
		return userSession, errors.New("user session not found")
	}
	return userSession, nil
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, accessToken string) error {
	return r.DB.Where("access_token = ?", accessToken).Delete(&models.UserSession{}).Error
}

func (r *UserRepository) UpdateTokenByRefreshToken(ctx context.Context, accessToken string, refreshToken string) error {
	return r.DB.Model(&models.UserSession{}).Where("refresh_token = ?", refreshToken).Update("access_token", accessToken).Error
}
