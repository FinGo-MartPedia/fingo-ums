package interfaces

import (
	"context"

	"github.com/fingo-martPedia/fingo-ums/internal/models"
	"github.com/gin-gonic/gin"
)

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}

type IRegisterHandler interface {
	Register(c *gin.Context)
}
