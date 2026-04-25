package adminusecase

import (
	"context"

	"github.com/ArthurTirta/monogo/pkg/dto"
)

type AdminUsecase interface {
	LoginAdmin(ctx context.Context, req *dto.ReqLogin) dto.ResAuthSingle
}
