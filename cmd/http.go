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
	healthcheckSvc := &services.Healthcheck{}
	healthcheckApi := &api.Healthcheck{
		HealthcheckService: healthcheckSvc,
	}

	registerRepo := &repository.RegisterRepository{
		DB: helpers.DB,
	}
	registerSvc := &services.RegisterService{
		RegisterRepository: registerRepo,
	}
	registerApi := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	r := gin.Default()

	r.GET("/health", healthcheckApi.Handler)

	apiV1 := r.Group("/api/v1/user")
	apiV1.POST("/register", registerApi.Register)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}
