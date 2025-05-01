package interfaces

import (
	"context"

	"github.com/fingo-martPedia/fingo-ums/internal/models/requests"
	"github.com/fingo-martPedia/fingo-ums/internal/models/responses"
)

type ILoginService interface {
	Login(ctx context.Context, request requests.LoginRequest) (responses.LoginResponse, error)
}
