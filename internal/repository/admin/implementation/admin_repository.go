package adminrepositoryimplementation

import (
	"context"
	"errors"

	"github.com/ArthurTirta/monogo/internal/entity"
	adminrepository "github.com/ArthurTirta/monogo/internal/repository/admin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type adminRepository struct {
	db    *gorm.DB
	admin entity.Admin
}

func NewAdminRepository(db *gorm.DB) adminrepository.AdminRepository {
	return &adminRepository{db: db, admin: entity.Admin{}}
}

func (r *adminRepository) Create(ctx context.Context, admin *entity.Admin) (output *entity.Admin, err error) {
	if r.db == nil {
		return nil, errors.New("database connection is not initialized")
	}
	if err := r.db.WithContext(ctx).Create(admin).Error; err != nil {
		return nil, err
	}
	return r.GetByID(ctx, admin.ID)
}

func (r *adminRepository) GetByID(ctx context.Context, id uuid.UUID) (output *entity.Admin, err error) {
	if r.db == nil {
		return nil, errors.New("database connection is not initialized")
	}
	if err := r.db.WithContext(ctx).Table(r.admin.TableName()).First(&output, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return
}

func (r *adminRepository) GetByEmail(ctx context.Context, email string) (output *entity.Admin, err error) {
	if r.db == nil {
		return nil, errors.New("database connection is not initialized")
	}
	if err := r.db.WithContext(ctx).Table(r.admin.TableName()).First(&output, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return
}

func (r *adminRepository) Update(ctx context.Context, id uuid.UUID, updateMap map[string]any) (output *entity.Admin, err error) {
	if r.db == nil {
		return nil, errors.New("database connection is not initialized")
	}
	if err := r.db.WithContext(ctx).Model(&r.admin).Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return nil, err
	}
	return r.GetByID(ctx, id)
}

func (r *adminRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	if r.db == nil {
		return errors.New("database connection is not initialized")
	}
	err = r.db.WithContext(ctx).Where("id = ?", id).Delete(&r.admin).Error
	return err
}
