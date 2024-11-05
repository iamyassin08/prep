package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/db"
)

// UserReq is the request struct for user creation and updates
type UserReq struct {
	UserID    string  `json:"UserID"`
	FirstName string  `json:"FirstName"`
	LastName  float64 `json:"LastName"`
	Email     float64 `json:"Email"`
}

type UserRes struct {
	db.User
	ExternalDetails UserExternalDetails `json:"ExternalDetails"`
}

//	@BasePath		/api/v1
//
// GetUser godoc
//
//	@Param			id	path	int	true	"User ID"
//
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
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	user, err := db.DB.GetUser(c.Context(), int32(iD))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	var res UserRes
	userCategory, _ := db.DB.GetUserCategory(c.Context(), user.ID)

	// Populating response structure
	res = UserRes{
		User: user,
		ExternalDetails: UserExternalDetails{
			UserID:        user.ID.String(),
			UserFirstName: user.FirstName,
			ExternalURL:   user.ExternalURL, // Assuming an ExternalURL exists on user
		},
	}

	return c.Status(fiber.StatusOK).JSON(res)
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
	users, err := db.DB.ListUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	userRes := []UserRes{}
	for _, user := range users {
		userCategory, _ := db.DB.GetUserCategory(c.Context(), user.ID)
		userRes = append(userRes, UserRes{
			User: user,
			ExternalDetails: UserExternalDetails{
				UserID:        user.ID.String(),
				UserFirstName: user.FirstName,
				ExternalURL:   user.ExternalURL,
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(userRes)
}

//	@BasePath		/api/v1
//
// CreateUser godoc
//
//	@Param			id	body	UserReq	true	"Create User Params"
//
//	@Summary		Create A New User
//	@Description	Create a new user by providing user data
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string
//	@Router			/users [post]
func (h *ApiHandler) CreateUser(c *fiber.Ctx) error {
	var userReq UserReq
	if err := c.BodyParser(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	userArg, err := createUserHelper(userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = db.DB.CreateUser(c.Context(), userArg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to create user")
	}
	return c.Status(fiber.StatusOK).JSON("User Created")
}

//	@BasePath		/api/v1
//
// UpdateUser godoc
//
//	@Param			id	path	int	true	"User ID"
//
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	false	"Authorization"
//	@Param			id	body	UserReq	true	"Update User Params"
//
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
		return c.Status(fiber.StatusInternalServerError).JSON("could not find user with ID: ", userId)
	}
	var userReq UserReq
	if err := c.BodyParser(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, "Could not convert JSON body to a struct")
	}
	userArg, err := updateUserHelper(userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userArg.ID = int32(iD)
	err = db.DB.UpdateUser(c.Context(), userArg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to update user with ID: " + userId + " Got Error: " + err.Error())
	}
	return c.Status(fiber.StatusOK).JSON("User Updated")
}

//	@BasePath		/api/v1
//
// DeleteUser godoc
//
//	@Param			id	path	int	true	"User ID"
//
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	false	"Authorization"
//	@Summary		Delete User By ID
//	@Description	Delete a user by providing user ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string
//	@Router			/users/{id} [delete]
func (h *ApiHandler) DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	iD, err := strconv.Atoi(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Invalid user ID")
	}
	err = db.DB.DeleteUser(c.Context(), int32(iD))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to delete user")
	}
	return c.Status(fiber.StatusOK).JSON("User Deleted")
}

// Helper Functions for DB Operations

func createUserHelper(user UserReq) (db.CreateUserParams, error) {
	var params db.CreateUserParams
	// Populate params for creation
	params.UserID = user.UserID
	params.FirstName = user.FirstName
	params.LastName = user.LastName
	params.Email = user.Email
	// You can add more fields as needed
	return params, nil
}

func updateUserHelper(user UserReq) (db.UpdateUserParams, error) {
	var params db.UpdateUserParams
	// Populate params for update
	params.UserID = user.UserID
	params.FirstName = user.FirstName
	params.LastName = user.LastName
	params.Email = user.Email
	return params, nil
}
