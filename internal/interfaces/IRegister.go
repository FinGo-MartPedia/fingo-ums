package interfaces

import (
	"context"

	"github.com/fingo-martPedia/fingo-ums/internal/models"
)

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}

type IRegisterRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
}
