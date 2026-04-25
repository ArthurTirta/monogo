package router

import (
	"github.com/ArthurTirta/monogo/internal/handler"
	userrepository "github.com/ArthurTirta/monogo/internal/repository/user/implementation"
	userserializer "github.com/ArthurTirta/monogo/internal/serializer/user/implementation"
	userusecase "github.com/ArthurTirta/monogo/internal/usecase/user/implementation"
)

func UserRouter(deps *Dependencies) {
	userRepository := userrepository.NewUserRepository(deps.DB)
	userSerializer := userserializer.NewUserSerializer()
	userUsecase := userusecase.NewUserUsecase(userRepository, userSerializer)
	userHandler := handler.NewUserHandler(userUsecase)

	userGroup := deps.App.Group("/v1/users")

	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Get("/:id", userHandler.GetUserByID)
	userGroup.Get("/", userHandler.GetUsersByFilter)
	userGroup.Put("/:id", userHandler.UpdateUser)
	userGroup.Delete("/:id", userHandler.DeleteUser)
}
