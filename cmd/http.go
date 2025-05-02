package cmd

import (
	"log"

	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/api"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/fingo-martPedia/fingo-ums/internal/repository"
	"github.com/fingo-martPedia/fingo-ums/internal/services"
	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependencies := dependencyInject()
	healthcheckApi := dependencies.HealthcheckAPI
	registerApi := dependencies.RegisterAPI
	loginApi := dependencies.LoginAPI
	logoutApi := dependencies.LogoutAPI

	r := gin.Default()

	r.GET("/health", healthcheckApi.Handler)

	apiV1 := r.Group("/api/v1/user")
	apiV1.POST("/register", registerApi.Register)
	apiV1.POST("/login", loginApi.Login)
	apiV1.DELETE("/logout", dependencies.MiddlewareValidateAuth, logoutApi.Logout)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	UserRepository interfaces.IUserRepository

	HealthcheckAPI interfaces.IHealthcheckHandler
	RegisterAPI    interfaces.IRegisterHandler
	LoginAPI       interfaces.ILoginHandler
	LogoutAPI      interfaces.ILogoutHandler
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

	logoutSvc := &services.LogoutService{
		UserRepository: userRepo,
	}
	logoutAPI := &api.LogoutHandler{
		LogoutService: logoutSvc,
	}

	return Dependency{
		UserRepository: userRepo,
		HealthcheckAPI: healthcheckAPI,
		RegisterAPI:    registerAPI,
		LoginAPI:       loginAPI,
		LogoutAPI:      logoutAPI,
	}
}
