package pasarrepositoryimplementation

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ArthurTirta/monogo/internal/entity"
	pasarrepository "github.com/ArthurTirta/monogo/internal/repository/pasar"
	"gorm.io/gorm"
)

type pasarRepository struct {
	db    *gorm.DB
	pasar entity.Pasar
}

func NewPasarRepository(db *gorm.DB) pasarrepository.PasarRepository {
	return &pasarRepository{db: db, pasar: entity.Pasar{}}
}

func (r *pasarRepository) Create(ctx context.Context, pasar *entity.Pasar) (output *entity.Pasar, err error) {
	if r.db == nil {
		return nil, errors.New("database connection is not initialized")
	}

	// Use a transaction to atomically determine next psr id and insert.
	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if pasar.ID == "" {
			var max sql.NullInt64
			// Extract numeric suffix from id like 'psr-001'
			q := fmt.Sprintf("SELECT COALESCE(MAX(CAST(SUBSTRING(id from 5) AS INTEGER)),0) as max FROM %s WHERE id LIKE 'psr-%%'", r.pasar.TableName())
			if err := tx.Raw(q).Scan(&max).Error; err != nil {
				return err
			}
			next := max.Int64 + 1
			pasar.ID = fmt.Sprintf("psr-%03d", next)
		}

		if err := tx.Table(r.pasar.TableName()).Create(pasar).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return pasar, nil
}

func (r *pasarRepository) GetAll(ctx context.Context) (output []entity.Pasar, err error) {
	if r.db == nil {
		return nil, errors.New("database connection is not initialized")
	}
	if err = r.db.WithContext(ctx).Table(r.pasar.TableName()).Find(&output).Error; err != nil {
		return nil, err
	}
	return output, nil
}
