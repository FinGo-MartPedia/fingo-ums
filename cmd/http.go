package cmd

import (
	"log"

	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/api"
	"github.com/fingo-martPedia/fingo-ums/internal/repository"
	"github.com/fingo-martPedia/fingo-ums/internal/services"
	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	Dependency := dependencyInject()
	healthcheckApi := Dependency.HealthcheckAPI
	registerApi := Dependency.RegisterAPI
	loginApi := Dependency.LoginAPI

	r := gin.Default()

	r.GET("/health", healthcheckApi.Handler)

	apiV1 := r.Group("/api/v1/user")
	apiV1.POST("/register", registerApi.Register)
	apiV1.POST("/login", loginApi.Login)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealthcheckAPI *api.Healthcheck
	RegisterAPI    *api.RegisterHandler
	LoginAPI       *api.LoginHandler
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckService: healthcheckSvc,
	}

	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	registerSvc := &services.RegisterService{
		UserRepository: userRepo,
	}
	registerAPI := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	loginSvc := &services.LoginService{
		UserRepository: userRepo,
	}
	loginAPI := &api.LoginHandler{
		LoginService: loginSvc,
	}

	return Dependency{
		HealthcheckAPI: healthcheckAPI,
		RegisterAPI:    registerAPI,
		LoginAPI:       loginAPI,
	}
}
