package router

import (
	"github.com/ArthurTirta/monogo/internal/handler"
	adminrepository "github.com/ArthurTirta/monogo/internal/repository/admin/implementation"
	adminusecase "github.com/ArthurTirta/monogo/internal/usecase/admin/implementation"
)

func AuthRouter(deps *Dependencies) {
	adminRepo := adminrepository.NewAdminRepository(deps.DB)
	adminUsecase := adminusecase.NewAdminUsecase(adminRepo, deps.Cfg)
	authHandler := handler.NewAuthHandler(adminUsecase)

	authGroup := deps.App.Group("/v1/auth")
	authGroup.Post("/login", authHandler.Login)
}
