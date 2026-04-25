package adminrepository

import (
	"context"

	"github.com/ArthurTirta/monogo/internal/entity"
	"github.com/google/uuid"
)

type AdminRepository interface {
	Create(ctx context.Context, admin *entity.Admin) (output *entity.Admin, err error)
	GetByID(ctx context.Context, id uuid.UUID) (output *entity.Admin, err error)
	GetByEmail(ctx context.Context, email string) (output *entity.Admin, err error)
	Update(ctx context.Context, id uuid.UUID, updateMap map[string]any) (output *entity.Admin, err error)
	Delete(ctx context.Context, id uuid.UUID) (err error)
}
