package router

import (
	"github.com/ArthurTirta/monogo/internal/handler"
	pasarrepository "github.com/ArthurTirta/monogo/internal/repository/pasar/implementation"
	pasarserializer "github.com/ArthurTirta/monogo/internal/serializer/pasar/implementation"
	pasarusecase "github.com/ArthurTirta/monogo/internal/usecase/pasar/implementation"
)

func PasarRouter(deps *Dependencies) {
	pasarRepo := pasarrepository.NewPasarRepository(deps.DB)
	pasarSerializer := pasarserializer.NewPasarSerializer()
	pasarUsecase := pasarusecase.NewPasarUsecase(pasarRepo, pasarSerializer)
	pasarHandler := handler.NewPasarHandler(pasarUsecase)

	pasarGroup := deps.App.Group("/v1/pasar")
	pasarGroup.Get("/", pasarHandler.GetPasarList)
	pasarGroup.Post("/", pasarHandler.CreatePasar)
}
