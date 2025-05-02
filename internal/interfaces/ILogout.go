package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ILogoutService interface {
	Logout(ctx context.Context, accessToken string) error
}

type ILogoutHandler interface {
	Logout(c *gin.Context)
}
