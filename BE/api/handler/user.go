package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/db"
)

type UserReq struct {
	UserID    string `json:"UserID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}

type UserRes struct {
	db.User
}

// ApiHandler holds dependencies for the handler
type ApiHandler struct {
	queries *db.Queries
}

// NewApiHandler creates a new API handler with the given dependencies
func NewApiHandler(queries *db.Queries) *ApiHandler {
	return &ApiHandler{
		queries: queries,
	}
}

// Populate user parameters for creation and update
func populateUserParams(user UserReq, params interface{}) error {
	switch p := params.(type) {
	case *db.CreateUserParams:
		p.FirstName = user.FirstName
		p.LastName = user.LastName
		p.Email = user.Email
	case *db.UpdateUserParams:
		p.FirstName = user.FirstName
		p.LastName = user.LastName
		p.Email = user.Email
	default:
		return fmt.Errorf("unsupported params type")
	}
	return nil
}

// Helper functions for creating and updating users
func createUserHelper(user UserReq) (db.CreateUserParams, error) {
	var params db.CreateUserParams
	err := populateUserParams(user, &params)
	return params, err
}

func updateUserHelper(user UserReq) (db.UpdateUserParams, error) {
	var params db.UpdateUserParams
	err := populateUserParams(user, &params)
	return params, err
}

//	@BasePath		/api/v1
//
// GetUser godoc
//
//	@Param			id	path	int	true	"User ID"
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	false	"Authorization"
//	@Summary		Get A User By ID
//	@Description	get User by Id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	UserRes
//	@Router			/users/{id} [get]
func (h *ApiHandler) ServeUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	iD, err := strconv.Atoi(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := h.queries.GetUser(c.Context(), int64(iD))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(UserRes{User: user})
}

//	@BasePath		/api/v1
//
// ListUsers godoc
//
//	@Summary		List Users
//	@Description	Get All Users in DB
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	UserRes
//	@Router			/users [get]
func (h *ApiHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.queries.ListUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	userRes := make([]UserRes, len(users))
	for i, user := range users {
		userRes[i] = UserRes{User: user}
	}

	return c.Status(fiber.StatusOK).JSON(userRes)
}

//	@BasePath		/api/v1
//
// UpdateUser godoc
//
//	@Param			id	path	int	true	"User ID"
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	false	"Authorization"
//	@Param			id				body	UserReq	true	"Update User Params"
//	@Summary		Update A User
//	@Description	Update User by Id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Router			/users/{id} [patch]
func (h *ApiHandler) UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	iD, err := strconv.Atoi(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var userReq UserReq
	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Could not convert JSON body to a struct"})
	}

	userArg, err := updateUserHelper(userReq)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userArg.ID = int64(iD)
	err = h.queries.UpdateUser(c.Context(), userArg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.Status(fiber.StatusOK).JSON("User Updated")
}

//	@BasePath		/api/v1
//
// DeleteUser godoc
//
//	@Param			id	path	int	true	"User ID"
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	false	"Authorization"
//	@Summary		Delete A User
//	@Description	Delete User by Id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Router			/users/{id} [delete]
func (h *ApiHandler) DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	iD, err := strconv.Atoi(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	err = h.queries.DeleteUser(c.Context(), int64(iD))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.Status(fiber.StatusOK).JSON("User Deleted")
}

//	@BasePath		/api/v1
//
// CreateUser godoc
//
//	@Param			user	body	UserReq	true	"Create User Params"
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	false	"Authorization"
//	@Summary		Create A User
//	@Description	Create a new user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	UserRes
//	@Router			/users [post]
func (h *ApiHandler) CreateUser(c *fiber.Ctx) error {
	var userReq UserReq
	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Could not parse JSON body"})
	}

	userArg, err := createUserHelper(userReq)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newUser, err := h.queries.CreateUser(c.Context(), userArg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(UserRes{User: newUser})
}
