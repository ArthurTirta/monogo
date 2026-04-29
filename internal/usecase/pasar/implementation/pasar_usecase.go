package pasarusecaseimplementation

import (
	"context"
	"net/http"

	"log/slog"

	pasarrepository "github.com/ArthurTirta/monogo/internal/repository/pasar"
	pasarserializer "github.com/ArthurTirta/monogo/internal/serializer/pasar"
	pasarusecase "github.com/ArthurTirta/monogo/internal/usecase/pasar"
	"github.com/ArthurTirta/monogo/pkg/dto"
	dtobase "github.com/ArthurTirta/monogo/pkg/dto/base"
	"github.com/go-playground/validator/v10"
)

type pasarUsecase struct {
	pasarRepository pasarrepository.PasarRepository
	serializer      pasarserializer.PasarSerializer
	validator       *validator.Validate
	logger          *slog.Logger
}

func NewPasarUsecase(repo pasarrepository.PasarRepository, serializer pasarserializer.PasarSerializer) pasarusecase.PasarUsecase {
	return &pasarUsecase{pasarRepository: repo, serializer: serializer, validator: validator.New(), logger: slog.Default().With("usecase", "pasar")}
}

func (u *pasarUsecase) CreatePasar(ctx context.Context, req *dto.ReqCreatePasar) dto.ResPasarSingle {
	// validate
	if req == nil {
		return dto.ResPasarSingleFromEntity(nil, http.StatusBadRequest, "request is nil", nil)
	}
	if err := req.Validate(u.validator); err != nil {
		return dto.ResPasarSingleFromEntity(nil, http.StatusBadRequest, err.Error(), nil)
	}

	ent, _ := u.serializer.CreateDTOToEntity(*req)
	created, err := u.pasarRepository.Create(ctx, &ent)
	if err != nil {
		s := err.Error()
		return dto.ResPasarSingleFromEntity(nil, http.StatusInternalServerError, "failed to create pasar", &s)
	}

	res := u.serializer.EntityToResponse(*created)
	return dto.ResPasarSingleFromEntity(&res, http.StatusCreated, "created", nil)
}

func (u *pasarUsecase) GetPasarList(ctx context.Context) dto.ResPasarList {
	items, err := u.pasarRepository.GetAll(ctx)
	if err != nil {
		s := err.Error()
		return dto.ResPasarList{BaseRes: dtobase.BaseRes{Success: false, Code: http.StatusInternalServerError, Message: "failed to fetch pasar", Stacktrace: &s}, Data: nil}
	}
	data := u.serializer.EntityListToResponse(items)
	return dto.ResPasarList{BaseRes: dtobase.BaseRes{Success: true, Code: http.StatusOK, Message: "ok"}, Data: data}
}
