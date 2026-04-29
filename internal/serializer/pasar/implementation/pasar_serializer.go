package pasarserializerimplementation

import (
	"github.com/ArthurTirta/monogo/internal/entity"
	pasarserializer "github.com/ArthurTirta/monogo/internal/serializer/pasar"
	"github.com/ArthurTirta/monogo/pkg/dto"
)

type pasarSerializer struct{}

func NewPasarSerializer() pasarserializer.PasarSerializer {
	return &pasarSerializer{}
}

func (s *pasarSerializer) CreateDTOToEntity(create dto.ReqCreatePasar) (entity.Pasar, error) {
	var out entity.Pasar
	out.Nama = create.Nama
	out.Longitude = create.Longitude
	out.Latitude = create.Latitude
	out.Alamat = create.Alamat
	if create.IsActive != nil {
		out.IsActive = *create.IsActive
	} else {
		out.IsActive = 1
	}
	return out, nil
}

func (s *pasarSerializer) EntityToResponse(e entity.Pasar) dto.ResPasar {
	return dto.ResPasar{
		ID:        e.ID,
		Nama:      e.Nama,
		Longitude: e.Longitude,
		Latitude:  e.Latitude,
		Alamat:    e.Alamat,
		IsActive:  e.IsActive,
	}
}

func (s *pasarSerializer) EntityListToResponse(es []entity.Pasar) []dto.ResPasar {
	out := make([]dto.ResPasar, len(es))
	for i, e := range es {
		out[i] = s.EntityToResponse(e)
	}
	return out
}
