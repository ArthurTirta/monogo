package pasarserializer

import (
	"github.com/ArthurTirta/monogo/internal/entity"
	"github.com/ArthurTirta/monogo/pkg/dto"
)

type PasarSerializer interface {
	CreateDTOToEntity(create dto.ReqCreatePasar) (entity.Pasar, error)
	EntityToResponse(e entity.Pasar) dto.ResPasar
	EntityListToResponse(es []entity.Pasar) []dto.ResPasar
}
