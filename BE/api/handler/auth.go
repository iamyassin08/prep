package handler

import (
	"github.com/gofiber/fiber/v2"
)

type ApiHandler struct {
}

//	@BasePath	/api/v1

// Health godoc
//
//	@Summary		Get API Status
//	@Description	Useful to preform a health-check
//	@Tags			Status
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Health	"One small step..."
//	@Router			/healthz [post]
func HealthCheck(c *fiber.Ctx) error {
	return c.Send([]byte("One small step...."))
}
