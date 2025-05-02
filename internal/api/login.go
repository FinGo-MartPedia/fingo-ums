package api

import (
	"errors"
	"net/http"

	"github.com/fingo-martPedia/fingo-ums/constants"
	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/fingo-martPedia/fingo-ums/internal/models/requests"
	"github.com/fingo-martPedia/fingo-ums/internal/models/responses"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (api *LoginHandler) Login(c *gin.Context) {
	log := helpers.Logger
	var req requests.LoginRequest
	var response responses.LoginResponse

	userAgent := c.Request.UserAgent()
	ipAddress := c.ClientIP()

	if userAgent == "" || ipAddress == "" {
		log.Error("User agent or IP address is empty")
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, "User agent or IP address is empty")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	req.UserAgent = userAgent
	req.IPAddress = ipAddress

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	response, err := api.LoginService.Login(c, req)
	if err != nil {
		log.Error("Failed to login user: ", err)
		if errors.Is(err, helpers.ErrInvalidCredentials) {
			helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		} else {
			helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrFailedServerError, err.Error())
		}
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, response)
}
