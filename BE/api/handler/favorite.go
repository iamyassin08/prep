package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/db"
)

//	@BasePath		/api/v1
//
// # FavoriteAProduct godoc
//
//	@Summary		Favorite A Product for user
//	@Param			id	path	string	true	"User ID"
//	@Param			productId	path	string	true	"Product ID"
//	@Description	Favorite A Product Using Product and User ID
//	@Tags			User
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	string
//	@Router			/user/{id}/favorites/{productId} [post]
func (h *ApiHandler) AddProductToFavorites(c *fiber.Ctx) error {
	profile, err := getUserProfileHelper(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	id, err := strconv.Atoi(c.Params("productId"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not parse product ID")
	}
	product, err := getProductHelper(c, int32(id))
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}
	err = db.DB.AddProductFavorite(c.Context(), db.AddProductFavoriteParams{
		ProductID: product.ID,
		ProfileID: profile.ID,
	})

	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}
	// favorites, err := db.DB.AddProductFavorite()
	return c.Status(fiber.StatusOK).JSON("Added Product to Favorites")
}

//	@BasePath		/api/v1
//
// GetFavoritesList godoc
//
//	@Summary		List Products Favorited By User
//	@Param			id	path	string	true	"User ID"
//	@Description	Get A List Of All Products that have been favorited by the user
//	@Tags			User
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	string
//	@Router			/user/{id}/favorites [get]
func (h *ApiHandler) GetUserFavorites(c *fiber.Ctx) error {
	profile, err := getUserProfileHelper(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	products, err := db.DB.GetUserFavorites(c.Context(), profile.ID)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}
	productRes := []ProductRes{}
	for _, product := range products {
		thumbnail, _ := db.DB.GetProductThumbnail(c.Context(), product.ID)
		imgs, _ := db.DB.GetProductImages(c.Context(), product.ID)
		productCategory, _ := db.DB.GetProductCategory(c.Context(), product.ID)
		if len(imgs) > 0 {
			img, _ := db.DB.GetProductThumbnailImage(c.Context(), thumbnail.ProductImageID)
			productRes = append(productRes, ProductRes{
				Product:      product,
				ThumbnailURL: img.ImageUrl,
				Category: ProductCategoryRes{
					ID:   productCategory.ID,
					Name: productCategory.Name,
				},
				Images: imgs})
		} else {
			productRes = append(productRes, ProductRes{
				Category: ProductCategoryRes{
					ID:   productCategory.ID,
					Name: productCategory.Name,
				},
				Product: product,
			})
		}
	}

	// favorites, err := db.DB.AddProductFavorite()
	return c.Status(fiber.StatusOK).JSON(productRes)
}

//	@BasePath		/api/v1
//
// RemoveProductFromFavorites godoc
//
//	@Summary		Remove Product from Favorites
//	@Param			id	path	string	true	"User ID"
//	@Param			productId	path	string	true	"Product ID"
//	@Description	Remove a product from a user's favorites list
//	@Tags			User
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	string
//	@Router			/user/{id}/favorites/{productId} [delete]
func (h *ApiHandler) RemoveProductToFavorites(c *fiber.Ctx) error {
	profile, err := getUserProfileHelper(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	id, err := strconv.Atoi(c.Params("productId"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not parse product ID")
	}
	product, err := getProductHelper(c, int32(id))
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}
	err = db.DB.RemoveProductfavorite(c.Context(), db.RemoveProductfavoriteParams{
		ProductID: product.ID,
		ProfileID: profile.ID,
	})
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}
	// favorites, err := db.DB.AddProductFavorite()
	return c.Status(fiber.StatusOK).JSON("Favorited A Product with ID: ")
}
