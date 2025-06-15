//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/api"
	"github.com/fingo-martPedia/fingo-ums/internal/api/grpcapi"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces"
	"github.com/fingo-martPedia/fingo-ums/internal/repository"
	"github.com/fingo-martPedia/fingo-ums/internal/services"
	"github.com/fingo-martPedia/fingo-ums/internal/services/grpcservice"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type Dependency struct {
	UserRepository interfaces.IUserRepository

	HealthcheckAPI  interfaces.IHealthcheckHandler
	LoginAPI        interfaces.ILoginHandler
	RegisterAPI     interfaces.IRegisterHandler
	LogoutAPI       interfaces.ILogoutHandler
	RefreshTokenAPI interfaces.IRefreshTokenHandler

	TokenValidationAPI *grpcapi.TokenValidationHandler
}

func provideDB() *gorm.DB {
	return helpers.DB
}

func InitDependency() Dependency {
	wire.Build(
		provideDB,

		repository.NewUserRepository,
		wire.Bind(new(interfaces.IUserRepository), new(*repository.UserRepository)),

		// === Services + Handlers + Interface Binding ===

		// Healthcheck
		services.NewHealthcheckService,
		api.NewHealthcheckHandler,
		wire.Bind(new(interfaces.IHealthcheckHandler), new(*api.Healthcheck)),

		// Login
		services.NewLoginService,
		api.NewLoginHandler,
		wire.Bind(new(interfaces.ILoginHandler), new(*api.LoginHandler)),

		// Register
		services.NewRegisterService,
		api.NewRegisterHandler,
		wire.Bind(new(interfaces.IRegisterHandler), new(*api.RegisterHandler)),

		// Logout
		services.NewLogoutService,
		api.NewLogoutHandler,
		wire.Bind(new(interfaces.ILogoutHandler), new(*api.LogoutHandler)),

		// RefreshToken
		services.NewRefreshTokenService,
		api.NewRefreshTokenHandler,
		wire.Bind(new(interfaces.IRefreshTokenHandler), new(*api.RefreshTokenHandler)),

		// TokenValidation
		grpcservice.NewTokenValidationService,
		grpcapi.NewTokenValidationHandler,

		// Final struct
		wire.Struct(new(Dependency), "*"),
	)
	return Dependency{}
}
