package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/db"
	"github.com/jackc/pgx/v5"
)

func (h ApiHandler) HandleGETSignup(c *fiber.Ctx) {
	c.JSON(fiber.StatusOK)
}

func (h ApiHandler) HandlePOSTLogout(c *fiber.Ctx) error {
	c.JSON(fiber.StatusOK)
	return nil
}

func (h ApiHandler) HandlePOSTLogin(c *fiber.Ctx) error {
	c.JSON(fiber.StatusOK)
	return nil
}

func getUserProfileHelper(c context.Context, profileID string) (db.Profile, error) {
	profile, err := db.DB.GetUserProfile(c, profileID)
	if err == pgx.ErrNoRows {
		err = db.DB.AddUserProfile(c, profileID)
		if err != nil {
			return db.Profile{}, err
		}
	} else if err != nil {
		return db.Profile{}, err
	}
	return profile, nil
}
