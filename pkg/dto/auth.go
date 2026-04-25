package dto

import (
	dtobase "github.com/ArthurTirta/monogo/pkg/dto/base"
	"github.com/go-playground/validator/v10"
)

type ReqLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *ReqLogin) Validate(validate *validator.Validate) error {
	return validate.Struct(r)
}

type ResAuthToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type ResAuthData struct {
	User  ResUser      `json:"user"`
	Token ResAuthToken `json:"token"`
}

type ResAuthSingle struct {
	dtobase.BaseRes
	Data *ResAuthData `json:"data"`
}
