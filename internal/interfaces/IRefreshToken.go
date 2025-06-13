package interfaces

import (
	"context"

	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/models/responses"
	"github.com/gin-gonic/gin"
)

type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (responses.RefreshTokenResponse, error)
}

type IRefreshTokenHandler interface {
	RefreshToken(c *gin.Context)
}
