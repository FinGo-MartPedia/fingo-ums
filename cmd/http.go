package cmd

import (
	"log"

	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependencies := InitDependency()

	r := gin.Default()

	r.GET("/health", dependencies.HealthcheckAPI.Handler)

	apiV1 := r.Group("/api/v1/user")
	apiV1.POST("/register", dependencies.RegisterAPI.Register)
	apiV1.POST("/login", dependencies.LoginAPI.Login)

	apiV1.DELETE("/logout", dependencies.MiddlewareValidateAuth, dependencies.LogoutAPI.Logout)
	apiV1.PUT("/refresh-token", dependencies.MiddlewareRefreshToken, dependencies.RefreshTokenAPI.RefreshToken)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}
