package handler

import (
	adminusecase "github.com/ArthurTirta/monogo/internal/usecase/admin"
	"github.com/ArthurTirta/monogo/pkg/dto"
	errorhelper "github.com/ArthurTirta/monogo/pkg/helper/error"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	adminUsecase adminusecase.AdminUsecase
}

func NewAuthHandler(adminUsecase adminusecase.AdminUsecase) *authHandler {
	return &authHandler{adminUsecase: adminUsecase}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var req dto.ReqLogin
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "code": fiber.StatusBadRequest, "message": err.Error(), "stacktrace": errorhelper.ComposeStacktrace(err)})
	}

	res := h.adminUsecase.LoginAdmin(c.Context(), &req)
	return c.Status(res.Code).JSON(res)
}
