package api

import (
	"net/http"

	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/fingo-martPedia/fingo-ums/internal/services"
	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	HealthcheckService interfaces.IHealthcheckServices
}

func NewHealthcheckHandler(svc *services.Healthcheck) *Healthcheck {
	return &Healthcheck{HealthcheckService: svc}
}

func (api *Healthcheck) Handler(c *gin.Context) {
	msg, err := api.HealthcheckService.HealthcheckServices()
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	helpers.SendResponse(c, http.StatusOK, msg, nil)
}
