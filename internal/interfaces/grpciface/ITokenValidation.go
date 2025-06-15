package grpciface

import (
	"context"

	"github.com/fingo-martPedia/fingo-ums/cmd/proto/tokenvalidation"
	"github.com/fingo-martPedia/fingo-ums/helpers"
)

type ITokenValidationHandler interface {
	TokenValidationHandler(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error)
}

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}
