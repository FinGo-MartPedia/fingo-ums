package repository

import (
	"context"

	"github.com/fingo-martPedia/fingo-ums/internal/models"
	"gorm.io/gorm"
)

type RegisterRepository struct {
	DB *gorm.DB
}

func (r *RegisterRepository) InsertNewUser(ctx context.Context, user *models.User) error {
	return r.DB.Create(&user).Error
}
