package api

import (
	"errors"
	"net/http"

	"github.com/fingo-martPedia/fingo-ums/constants"
	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/fingo-martPedia/fingo-ums/internal/models"
	"github.com/fingo-martPedia/fingo-ums/internal/services"
	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	RegisterService interfaces.IRegisterService
}

func NewRegisterHandler(svc *services.RegisterService) *RegisterHandler {
	return &RegisterHandler{RegisterService: svc}
}

func (api *RegisterHandler) Register(c *gin.Context) {
	log := helpers.Logger
	req := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	resp, err := api.RegisterService.Register(c, req)
	if err != nil {
		log.Error("Failed to register user: ", err)
		if errors.Is(err, helpers.ErrUsernameExists) {
			helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		} else {
			helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrFailedServerError, err.Error())
		}
		return
	}
	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, resp)
}
