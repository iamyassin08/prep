package middlewares

import (
	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt/v5"
	"github.com/iamyassin08/prep/shared"
)

func RequiresRealmRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		claims := ctx.Value(shared.ContextKeyClaims).(golangJwt.MapClaims)
		jwtHelper := shared.NewJwtHelper(claims)
		if !jwtHelper.IsUserInRealmRole(role) {
			return c.Status(fiber.StatusUnauthorized).SendString("role authorization failed")
		}
		return c.Next()
	}
}
