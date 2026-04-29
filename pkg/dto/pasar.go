package dto

import (
	"net/http"

	dtobase "github.com/ArthurTirta/monogo/pkg/dto/base"
	"github.com/go-playground/validator/v10"
)

type ReqCreatePasar struct {
	Nama      string  `json:"nama" validate:"required"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Alamat    string  `json:"alamat"`
	IsActive  *int    `json:"is_active"`
}

func (r *ReqCreatePasar) Validate(validate *validator.Validate) error {
	return validate.Struct(r)
}

type ResPasar struct {
	ID        string  `json:"id"`
	Nama      string  `json:"nama"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Alamat    string  `json:"alamat"`
	IsActive  int     `json:"is_active"`
}

type ResPasarSingle struct {
	dtobase.BaseRes
	Data *ResPasar `json:"data"`
}

type ResPasarList struct {
	dtobase.BaseRes
	Data []ResPasar `json:"data"`
}

func ResPasarSingleFromEntity(data *ResPasar, code int, message string, stacktrace *string) ResPasarSingle {
	isSuccess := code >= http.StatusOK && code < http.StatusMultipleChoices
	return ResPasarSingle{BaseRes: dtobase.BaseRes{Success: isSuccess, Code: code, Message: message, Stacktrace: stacktrace}, Data: data}
}
