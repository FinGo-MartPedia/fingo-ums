package api

import (
	"net/http"

	"github.com/fingo-martPedia/fingo-ums/constants"
	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/fingo-martPedia/fingo-ums/internal/services"
	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func NewRefreshTokenHandler(svc *services.RefreshTokenService) *RefreshTokenHandler {
	return &RefreshTokenHandler{RefreshTokenService: svc}
}

func (api *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	log := helpers.Logger

	refreshTokenRaw, _ := c.Get("refreshToken")
	refreshToken := refreshTokenRaw.(string)

	claimRaw, ok := c.Get("claim")
	if !ok {
		log.Error("Failed to get claim token from context")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrFailedServerError, nil)
		return
	}

	claimToken, ok := claimRaw.(*helpers.ClaimToken)
	if !ok {
		log.Error("Failed to parse claim raw to claim token")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrFailedServerError, nil)
		return
	}

	response, err := api.RefreshTokenService.RefreshToken(c.Request.Context(), refreshToken, *claimToken)
	if err != nil {
		log.Error("Failed to refresh token: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrFailedServerError, err.Error())
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, response)
}
