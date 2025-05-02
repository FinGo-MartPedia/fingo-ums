package api

import (
	"net/http"

	"github.com/fingo-martPedia/fingo-ums/constants"
	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (h *LogoutHandler) Logout(c *gin.Context) {
	log := helpers.Logger

	accessTokenRaw, _ := c.Get("accessToken")
	accessToken := accessTokenRaw.(string)

	err := h.LogoutService.Logout(c, accessToken)
	if err != nil {
		log.Error("Failed to logout user: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrFailedServerError, err.Error())
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, nil)
}
