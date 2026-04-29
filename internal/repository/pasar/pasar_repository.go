package pasarrepository

import (
	"context"

	"github.com/ArthurTirta/monogo/internal/entity"
)

type PasarRepository interface {
	Create(ctx context.Context, pasar *entity.Pasar) (output *entity.Pasar, err error)
	GetAll(ctx context.Context) (output []entity.Pasar, err error)
}
