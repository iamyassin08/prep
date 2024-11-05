package handler

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserExternalDetails struct {
	UserFirstName string `json:"UserFirstName"`
	UserID        string `json:"UserID"`
	UserImageURL  string `json:"UserImageURL"`
	ExternalURL   string `json:"ExternalURL"`
}

type UserReq struct {
	UserID    string  `json:"UserID"`
	FirstName string  `json:"FirstName"`
	LastName  float64 `json:"LastName "`
	Email     float64 `json:"Email "`
}
type UserRes struct {
	db.User
	ThumbnailURL    string              `json:"Thumbnail_url"`
	Images          []db.UserImage      `json:"Images"`
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
	imgs, _ := db.DB.GetUserImages(c.Context(), user.ID)
	thumbnail, _ := db.DB.GetUserThumbnail(c.Context(), user.ID)
	var res UserRes
	userCategory, _ := db.DB.GetUserCategory(c.Context(), user.ID)

	if len(imgs) > 0 {
		img, _ := db.DB.GetUserThumbnailImage(c.Context(), thumbnail.UserImageID)
		res = UserRes{
			User:         user,
			Images:       imgs,
			ThumbnailURL: img.ImageUrl,
			Category: UserCategoryRes{
				ID:        userCategory.ID,
				FirstName: userCategory.FirstName,
			},
		}
	} else {
		res = UserRes{
			Category: UserCategoryRes{
				ID:        userCategory.ID,
				FirstName: userCategory.FirstName,
			},
			User: user,
		}
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

//	@BasePath		/api/v1
//
// # ListUsers godoc
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
		thumbnail, _ := db.DB.GetUserThumbnail(c.Context(), user.ID)
		imgs, _ := db.DB.GetUserImages(c.Context(), user.ID)
		userCategory, _ := db.DB.GetUserCategory(c.Context(), user.ID)
		if len(imgs) > 0 {
			img, _ := db.DB.GetUserThumbnailImage(c.Context(), thumbnail.UserImageID)
			userRes = append(userRes, UserRes{
				User:         user,
				ThumbnailURL: img.ImageUrl,
				Category: UserCategoryRes{
					ID:        userCategory.ID,
					FirstName: userCategory.FirstName,
				},
				Images: imgs})
		} else {
			userRes = append(userRes, UserRes{
				User: user,
				Category: UserCategoryRes{
					ID:        userCategory.ID,
					FirstName: userCategory.FirstName,
				},
				Images: []db.UserImage{}})
		}
	}
	return c.Status(fiber.StatusOK).JSON(userRes)
}

//	@BasePath		/api/v1
//
// UpdateUser godoc
//
//	@Param			id	path	int	true	"User ID"
//
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	false	"Authorization"
//
// @Param			id	body	UserReq	true	"Update User Params"
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
// UpdateUser godoc
//
//	@Param			id	path	int	true	"User ID"
//
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	false	"Authorization"
//	@Summary		Get User's Thumbnail
//	@Description	Retrieve The User Image that was Set as the thumbnail
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	db.UserImage
//	@Router			/users/{id}/thumbnail [get]
func (h *ApiHandler) GetUserThumbnail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not parse user ID")
	}
	user, err := getUserHelper(c, int32(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not find user with ID: " + c.Params("id"))
	}
	thumbnail, _ := db.DB.GetUserThumbnail(c.Context(), user.ID)
	return c.Status(fiber.StatusOK).JSON(thumbnail)
}

func setThumbnailHelper(c *fiber.Ctx, id int32, imageID int32) error {
	old, err := db.DB.GetUserThumbnail(c.Context(), id)
	if err != nil && err != pgx.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	fmt.Printf("OLD THUMBNAIL ID: %d\n", old.UserImageID)
	if err == pgx.ErrNoRows {
		fmt.Println("SETTING IMAGE AS THUMBNAIL")
		db.DB.SetUserThumbnail(c.Context(), db.SetUserThumbnailParams{
			UserID:      id,
			UserImageID: imageID,
		})
	} else {
		fmt.Println("UPDATING PRODUCT'S EXISTING THUMBNAIL")
		err = db.DB.UpdateUserThumbnail(c.Context(), db.UpdateUserThumbnailParams{
			UserID:      id,
			UserImageID: imageID,
		})
		fmt.Printf("UPDATED THUMBNAIL ID: %d\n", imageID)
		if err != nil {
			return err
		}
	}
	return nil
}

func createUserImageHelper(c *fiber.Ctx, id int32, imageURL string) (db.UserImage, error) {
	existingImage, err := db.DB.GetUserImage(c.Context(), db.GetUserImageParams{
		UserID:   id,
		ImageUrl: imageURL,
	})
	if err == pgx.ErrNoRows {
		fmt.Println("CREATING NEW IMAGE")
		userImg, err := db.DB.CreateUserImage(c.Context(), db.CreateUserImageParams{
			UserID:   id,
			ImageUrl: imageURL,
		})
		if err != nil {
			return db.UserImage{}, err
		}
		return userImg, nil
	} else if err != nil {
		fmt.Println(err.Error())
		return db.UserImage{}, err
	}
	fmt.Println("IMAGE ALREADY EXISTS")
	return existingImage, nil
}

func getUserHelper(c *fiber.Ctx, userId int32) (db.User, error) {
	user, err := db.DB.GetUser(c.Context(), userId)
	if err != nil {
		return db.User{}, err
	}
	return user, err
}

func populateUserParams(user UserReq, params interface{}) error {
	var err error

	scanNumeric := func(dst *pgtype.Numeric, src float64) {
		num := fmt.Sprintf("%.2f", src)
		// strconv.FormatFloat(src, 'E', 2, 64)
		err := dst.Scan(num)
		if err != nil {
			fmt.Println("ERRRRRROOOOOORRRRR")
			fmt.Println(err.Error())
			dst.Int = big.NewInt(int64(src))
			dst.Exp = 0
			dst.Valid = true
		}
	}

	scanText := func(dst *pgtype.Text, src string) {
		if err == nil {
			dst.String = src
			dst.Valid = true
		}
	}

	switch p := params.(type) {
	case *db.CreateUserParams, *db.UpdateUserParams:
		var price, regEmail, disEmail *pgtype.Numeric
		var desc, shortDesc, pType, profileID *pgtype.Text
		var title *string
		var quantity *int32

		if createParams, ok := params.(*db.CreateUserParams); ok {
			price = &createParams.Email
			regEmail = &createParams.LastName
			disEmail = &createParams.DiscountEmail
			desc = &createParams.Description
			shortDesc = &createParams.ShortDescription
			pType = &createParams.Type
			profileID = &createParams.UserID
			title = &createParams.FirstName
			quantity = &createParams.Quantity
		} else if updateParams, ok := params.(*db.UpdateUserParams); ok {
			price = &updateParams.Email
			regEmail = &updateParams.LastName
			disEmail = &updateParams.DiscountEmail
			desc = &updateParams.Description
			shortDesc = &updateParams.ShortDescription
			pType = &updateParams.Type
			profileID = &updateParams.UserID
			title = &updateParams.FirstName
			quantity = &updateParams.Quantity
		}

		scanNumeric(price, user.Email)
		scanNumeric(regEmail, user.LastName)
		scanNumeric(disEmail, user.DiscountEmail)
		scanText(desc, user.Description)
		scanText(shortDesc, user.ShortDescription)
		scanText(pType, user.Type)
		scanText(profileID, user.UserID)
		*title = user.FirstName
		*quantity = user.Quantity
		fmt.Println(p)
	default:
		return fmt.Errorf("unsupported params type")
	}

	if err != nil {
		return err
	}

	switch user.Type {
	case PRODUCT_RESERVATION, PRODUCT_DELISTED, PRODUCT_EXTERNAL, PRODUCT_BUY:
	default:
		return fmt.Errorf("user has unknown 'Type'")
	}

	return nil
}

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
