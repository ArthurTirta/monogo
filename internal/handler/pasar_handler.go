package handler

import (
	"fmt"

	pasarusecase "github.com/ArthurTirta/monogo/internal/usecase/pasar"
	"github.com/ArthurTirta/monogo/pkg/dto"
	errorhelper "github.com/ArthurTirta/monogo/pkg/helper/error"
	"github.com/gofiber/fiber/v2"
)

type pasarHandler struct {
	pasarUsecase pasarusecase.PasarUsecase
}

// NewPasarHandler constructs a pasar handler with a PasarUsecase (hexagonal wiring)
func NewPasarHandler(u pasarusecase.PasarUsecase) *pasarHandler {
	return &pasarHandler{pasarUsecase: u}
}

// GetPasarList delegates to usecase to fetch pasar list
func (h *pasarHandler) GetPasarList(c *fiber.Ctx) error {
	res := h.pasarUsecase.GetPasarList(c.Context())
	return c.Status(res.Code).JSON(res)
}

// CreatePasar handles POST /v1/pasar
func (h *pasarHandler) CreatePasar(c *fiber.Ctx) error {
	// debug: log raw body for troubleshooting
	raw := c.Body()
	fmt.Printf("[DEBUG] CreatePasar raw body: %s\n", string(raw))

	var req dto.ReqCreatePasar
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":    false,
			"code":       fiber.StatusBadRequest,
			"message":    err.Error(),
			"stacktrace": errorhelper.ComposeStacktrace(err),
		})
	}

	res := h.pasarUsecase.CreatePasar(c.Context(), &req)
	return c.Status(res.Code).JSON(res)
}
