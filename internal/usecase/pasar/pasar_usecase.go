package pasarusecase

import (
	"context"

	"github.com/ArthurTirta/monogo/pkg/dto"
)

type PasarUsecase interface {
	CreatePasar(ctx context.Context, req *dto.ReqCreatePasar) dto.ResPasarSingle
	GetPasarList(ctx context.Context) dto.ResPasarList
}
