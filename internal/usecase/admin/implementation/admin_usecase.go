package adminusecaseimplementation

import (
	"context"
	"net/http"
	"time"

	"github.com/ArthurTirta/monogo/config"
	// "github.com/ArthurTirta/monogo/internal/entity"
	adminrepository "github.com/ArthurTirta/monogo/internal/repository/admin"
	adminusecase "github.com/ArthurTirta/monogo/internal/usecase/admin"
	"github.com/ArthurTirta/monogo/pkg/dto"
	dtobase "github.com/ArthurTirta/monogo/pkg/dto/base"
	errorhelper "github.com/ArthurTirta/monogo/pkg/helper/error"
	jwthelper "github.com/ArthurTirta/monogo/pkg/helper/jwt"
	passwordhelper "github.com/ArthurTirta/monogo/pkg/helper/password"
	"github.com/go-playground/validator/v10"
	// "github.com/google/uuid"
	"log/slog"
)

type adminUsecase struct {
	adminRepository adminrepository.AdminRepository
	cfg             *config.Config
	logger          *slog.Logger
	validator       *validator.Validate
}

func NewAdminUsecase(adminRepository adminrepository.AdminRepository, cfg *config.Config) adminusecase.AdminUsecase {
	return &adminUsecase{
		adminRepository: adminRepository,
		cfg:             cfg,
		logger:          slog.Default().With("usecase", "admin"),
		validator:       validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (u *adminUsecase) LoginAdmin(ctx context.Context, req *dto.ReqLogin) dto.ResAuthSingle {
	u.logger.InfoContext(ctx, "login attempt", "email", req.Email)
	select {
	case <-ctx.Done():
		u.logger.ErrorContext(ctx, "LoginAdmin: context done", "error", ctx.Err())
		return dto.ResAuthSingle{BaseRes: dtobase.BaseRes{Success: false, Code: http.StatusInternalServerError, Message: "internal error", Stacktrace: errorhelper.ComposeStacktrace(ctx.Err())}}
	default:
	}

	if req == nil {
		return dto.ResAuthSingle{BaseRes: dtobase.BaseRes{Success: false, Code: http.StatusBadRequest, Message: "request is nil", Stacktrace: errorhelper.ComposeStacktrace(nil)}}
	}

	if err := u.validator.Struct(req); err != nil {
		return dto.ResAuthSingle{BaseRes: dtobase.BaseRes{Success: false, Code: http.StatusBadRequest, Message: err.Error(), Stacktrace: errorhelper.ComposeStacktrace(err)}}
	}

	admin, err := u.adminRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		u.logger.ErrorContext(ctx, "LoginAdmin: admin not found or repo error", "email", req.Email, "error", err.Error())
		return dto.ResAuthSingle{BaseRes: dtobase.BaseRes{Success: false, Code: http.StatusUnauthorized, Message: "invalid credentials", Stacktrace: errorhelper.ComposeStacktrace(err)}}
	}

	if !passwordhelper.VerifyPassword(admin.Password, req.Password) {
		u.logger.ErrorContext(ctx, "LoginAdmin: invalid password", "email", req.Email)
		return dto.ResAuthSingle{BaseRes: dtobase.BaseRes{Success: false, Code: http.StatusUnauthorized, Message: "invalid credentials"}}
	}

	token, exp, err := jwthelper.GenerateAccessToken(&u.cfg.JWTConfig, admin.ID)
	if err != nil {
		u.logger.ErrorContext(ctx, "LoginAdmin: token generation failed", "error", err.Error())
		return dto.ResAuthSingle{BaseRes: dtobase.BaseRes{Success: false, Code: http.StatusInternalServerError, Message: "failed to generate token", Stacktrace: errorhelper.ComposeStacktrace(err)}}
	}

	expiresIn := int(time.Unix(exp, 0).Sub(time.Now()).Seconds())

	resUser := dto.ResUser{
		ID:       admin.ID,
		Name:     admin.Name,
		Email:    admin.Email,
		Status:   0,
		Metadata: dto.UserMetadata{},
	}

	data := &dto.ResAuthData{
		User: resUser,
		Token: dto.ResAuthToken{
			AccessToken: token,
			ExpiresIn:   expiresIn,
			TokenType:   "Bearer",
		},
	}

	return dto.ResAuthSingle{BaseRes: dtobase.BaseRes{Success: true, Code: http.StatusOK, Message: "login successful"}, Data: data}
}
