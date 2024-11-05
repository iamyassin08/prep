package handler

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/shagreen"
	"github.com/pkg/errors"
)

type ApiHandler struct {
}

type RegisterUseCase interface {
	Register(context.Context, shagreen.RegistrationRequest) (*shagreen.RegistrationResponse, error)
	Login(context.Context, shagreen.LoginRequest) (*shagreen.LoginResponse, error)
}

//	@BasePath	/api/v1

// Register godoc
//
//	@param			request	body shagreen.RegistrationRequest	true	"Registration Request"
//
// @Summary		Register using API
// @Description	Send cgreenentials to get token
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Success		200	{string}	UserToken
// @Router			/register [post]
func RegisterHandler(uc RegisterUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		var request = shagreen.RegistrationRequest{}

		err := c.BodyParser(&request)
		if err != nil {
			return errors.Wrap(err, "unable to parse incoming request")
		}
		fmt.Println(request)

		response, err := uc.Register(ctx, request)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusCreated).JSON(response)
	}
}

//	@BasePath	/api/v1

// Login godoc
//
//	@param			request	body shagreen.LoginRequest	true	"Login Request"
//	@Summary		Login using API
//	@Description	Send cgreenentials to get login token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	UserToken
//	@Router			/login [post]
func LoginHandler(uc RegisterUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		var request = shagreen.LoginRequest{}

		err := c.BodyParser(&request)
		if err != nil {
			return errors.Wrap(err, "unable to parse incoming request")
		}
		fmt.Println(request)
		response, err := uc.Login(ctx, request)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(response.ResToken.AccessToken)
	}
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
