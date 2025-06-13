package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		log.Println("authorization empty")
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		log.Println("invalid authorization format")
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	accessToken := authHeader[len(bearerPrefix):]

	user, err := d.UserRepository.GetUserSessionByAccessToken(ctx.Request.Context(), accessToken)
	if err != nil {
		log.Println("failed to get user session on DB: ", err)
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	claim, err := helpers.ValidateToken(ctx.Request.Context(), accessToken)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("accessToken", accessToken)
	ctx.Set("claim", claim)
	ctx.Set("userId", user.UserID)

	ctx.Next()
}

func (d *Dependency) MiddlewareRefreshToken(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		log.Println("authorization empty")
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		log.Println("invalid authorization format")
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	refreshToken := authHeader[len(bearerPrefix):]

	_, err := d.UserRepository.GetUserSessionByRefreshToken(ctx.Request.Context(), refreshToken)
	if err != nil {
		log.Println("failed to get user session on DB: ", err)
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	claim, err := helpers.ValidateToken(ctx.Request.Context(), refreshToken)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("refreshToken", refreshToken)
	ctx.Set("claim", claim)

	ctx.Next()
}
