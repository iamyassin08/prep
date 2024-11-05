package handler

import (
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProductExternalDetails struct {
	BrandName     string `json:"BrandName"`
	BrandID       string `json:"BrandID"`
	BrandImageURL string `json:"BrandImageURL"`
	ExternalURL   string `json:"ExternalURL"`
}

type ProductCategoryRes struct {
	Name string `json:"Name"`
	ID   int32  `json:"ID"`
}
type ProductReq struct {
	ProfileID        string  `json:"ProfileID"`
	Title            string  `json:"Title"`
	Price            float64 `json:"Price"`
	RegularPrice     float64 `json:"RegularPrice"`
	DiscountPrice    float64 `json:"DiscountPrice"`
	Quantity         int32   `json:"Quantity"`
	Description      string  `json:"Description"`
	ShortDescription string  `json:"ShortDescription"`
	Type             string  `json:"Type"`
}
type ProductRes struct {
	db.Product
	ThumbnailURL    string                 `json:"Thumbnail_url"`
	Images          []db.ProductImage      `json:"Images"`
	Category        ProductCategoryRes     `json:"Category"`
	ExternalDetails ProductExternalDetails `json:"ExternalDetails"`
}

//	@BasePath		/api/v1
//
// GetProduct godoc
//
//	@Param			id	path	int	true	"Product ID"
//
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	false	"Authorization"
//	@Summary		Get A Product By ID
//	@Description	get Product by Id
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ProductRes
//	@Router			/products/{id} [get]
func (h *ApiHandler) ServeProduct(c *fiber.Ctx) error {
	productId := c.Params("id")
	iD, err := strconv.Atoi(productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	product, err := db.DB.GetProduct(c.Context(), int32(iD))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	imgs, _ := db.DB.GetProductImages(c.Context(), product.ID)
	thumbnail, _ := db.DB.GetProductThumbnail(c.Context(), product.ID)
	var res ProductRes
	productCategory, _ := db.DB.GetProductCategory(c.Context(), product.ID)

	if len(imgs) > 0 {
		img, _ := db.DB.GetProductThumbnailImage(c.Context(), thumbnail.ProductImageID)
		res = ProductRes{
			Product:      product,
			Images:       imgs,
			ThumbnailURL: img.ImageUrl,
			Category: ProductCategoryRes{
				ID:   productCategory.ID,
				Name: productCategory.Name,
			},
		}
	} else {
		res = ProductRes{
			Category: ProductCategoryRes{
				ID:   productCategory.ID,
				Name: productCategory.Name,
			},
			Product: product,
		}
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

//	@BasePath		/api/v1
//
// # ListProducts godoc
//
//	@Summary		List Products
//	@Description	Get All Products in DB
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	ProductRes
//	@Router			/products [get]
func (h *ApiHandler) ListProducts(c *fiber.Ctx) error {
	products, err := db.DB.ListProducts(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
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
				Product: product,
				Category: ProductCategoryRes{
					ID:   productCategory.ID,
					Name: productCategory.Name,
				},
				Images: []db.ProductImage{}})
		}
	}
	return c.Status(fiber.StatusOK).JSON(productRes)
}

//	@BasePath		/api/v1
//
// UpdateProduct godoc
//
//	@Param			id	path	int	true	"Product ID"
//
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	false	"Authorization"
//
// @Param			id	body	ProductReq	true	"Update Product Params"
//
//	@Summary		Update A Product
//	@Description	Update Product by Id
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Router			/products/{id} [patch]
func (h *ApiHandler) UpdateProduct(c *fiber.Ctx) error {
	productId := c.Params("id")
	iD, err := strconv.Atoi(productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("could not find product with ID: ", productId)
	}
	var productReq ProductReq
	if err := c.BodyParser(&productReq); err != nil {
		return c.JSON(http.StatusBadRequest, "Could not convert JSON body to a struct")
	}
	productArg, err := updateProductHelper(productReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	productArg.ID = int32(iD)
	err = db.DB.UpdateProduct(c.Context(), productArg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to update product with ID: " + productId + " Got Error: " + err.Error())
	}
	return c.Status(fiber.StatusOK).JSON("Product Updated")
}

//	@BasePath		/api/v1
//
// UpdateProduct godoc
//
//	@Param			id	path	int	true	"Product ID"
//
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	false	"Authorization"
//	@Summary		Get Product's Thumbnail
//	@Description	Retrieve The Product Image that was Set as the thumbnail
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	db.ProductImage
//	@Router			/products/{id}/thumbnail [get]
func (h *ApiHandler) GetProductThumbnail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not parse product ID")
	}
	product, err := getProductHelper(c, int32(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not find product with ID: " + c.Params("id"))
	}
	thumbnail, _ := db.DB.GetProductThumbnail(c.Context(), product.ID)
	return c.Status(fiber.StatusOK).JSON(thumbnail)
}

func setThumbnailHelper(c *fiber.Ctx, id int32, imageID int32) error {
	old, err := db.DB.GetProductThumbnail(c.Context(), id)
	if err != nil && err != pgx.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	fmt.Printf("OLD THUMBNAIL ID: %d\n", old.ProductImageID)
	if err == pgx.ErrNoRows {
		fmt.Println("SETTING IMAGE AS THUMBNAIL")
		db.DB.SetProductThumbnail(c.Context(), db.SetProductThumbnailParams{
			ProductID:      id,
			ProductImageID: imageID,
		})
	} else {
		fmt.Println("UPDATING PRODUCT'S EXISTING THUMBNAIL")
		err = db.DB.UpdateProductThumbnail(c.Context(), db.UpdateProductThumbnailParams{
			ProductID:      id,
			ProductImageID: imageID,
		})
		fmt.Printf("UPDATED THUMBNAIL ID: %d\n", imageID)
		if err != nil {
			return err
		}
	}
	return nil
}

func createProductImageHelper(c *fiber.Ctx, id int32, imageURL string) (db.ProductImage, error) {
	existingImage, err := db.DB.GetProductImage(c.Context(), db.GetProductImageParams{
		ProductID: id,
		ImageUrl:  imageURL,
	})
	if err == pgx.ErrNoRows {
		fmt.Println("CREATING NEW IMAGE")
		productImg, err := db.DB.CreateProductImage(c.Context(), db.CreateProductImageParams{
			ProductID: id,
			ImageUrl:  imageURL,
		})
		if err != nil {
			return db.ProductImage{}, err
		}
		return productImg, nil
	} else if err != nil {
		fmt.Println(err.Error())
		return db.ProductImage{}, err
	}
	fmt.Println("IMAGE ALREADY EXISTS")
	return existingImage, nil
}

//	@BasePath		/api/v1
//
// UploadThumbnail godoc
//
//	@Param			id	path	int	true	"Product ID"
//
//	@Security		ApiKeyAuth
//	@param			Authorization	header		string	false	"Authorization"
//	@Param			image			formData	file	true	"image to upload"
//	@Summary		Upload Image And Set As Thumbnail Product
//	@Description	Update Product by Id
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	string
//	@failure		default	{object}	string
//	@Router			/products/{id}/thumbnail [post]
func (h *ApiHandler) UploadThumbnail(c *fiber.Ctx) error {
	bucketName := os.Getenv("MINIO_BUCKET")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not parse product ID")
	}
	product, err := getProductHelper(c, int32(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not find product with ID: " + c.Params("id"))
	}
	objectName, err := uploadFileHelper("products", product.ID, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	imageURL := fmt.Sprintf("https://%s/%s/%s", minioEndpoint, bucketName, objectName)
	productImg, err := createProductImageHelper(c, product.ID, imageURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not Create Image: ", err.Error())
	}
	err = setThumbnailHelper(c, product.ID, productImg.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	thumbnail, _ := db.DB.GetProductThumbnail(c.Context(), product.ID)
	img, _ := db.DB.GetProductThumbnailImage(c.Context(), thumbnail.ProductImageID)
	type ProductThumbnail struct {
		Thumbnail db.ProductImage `json:"Thumbnail"`
	}
	return c.Status(fiber.StatusOK).JSON(&ProductThumbnail{
		Thumbnail: img,
	})
}

//	@BasePath		/api/v1
//
// UploadThumbnail godoc
//
//	@Param			id	path	int	true	"Product ID"
//
//	@Security		ApiKeyAuth
//	@param			Authorization	header		string	false	"Authorization"
//	@Param			image			formData	file	true	"image to upload"
//	@Summary		Upload Image And Set As Thumbnail Product
//	@Description	Upload An Image for the Product with This ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	string
//	@failure		default	{object}	string
//	@Router			/products/{id}/images [post]
func (h *ApiHandler) UploadFile(c *fiber.Ctx) error {
	bucketName := os.Getenv("MINIO_BUCKET")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not parse product ID")
	}
	product, err := getProductHelper(c, int32(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not find product with ID: " + c.Params("id"))
	}
	objectName, err := uploadFileHelper("products", product.ID, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not Upload Image: " + err.Error())
	}
	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	imageURL := fmt.Sprintf("https://%s/%s/%s", minioEndpoint, bucketName, objectName)
	_, err = createProductImageHelper(c, product.ID, imageURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not Create Image: " + err.Error())
	}
	imgs, _ := db.DB.GetProductImages(c.Context(), product.ID)
	thumbnail, _ := db.DB.GetProductThumbnail(c.Context(), product.ID)
	img, _ := db.DB.GetProductThumbnailImage(c.Context(), thumbnail.ProductImageID)
	type ProductImages struct {
		Images    []db.ProductImage `json:"Images"`
		Thumbnail db.ProductImage   `json:"Thumbnail"`
	}
	return c.Status(fiber.StatusOK).JSON(&ProductImages{
		Images:    imgs,
		Thumbnail: img,
	})
}

//	@BasePath		/api/v1
//
// UploadThumbnail godoc
//
//	@Param			id	path	int	true	"Product ID"
//
//	@Security		ApiKeyAuth
//	@param			Authorization	header		string	false	"Authorization"
//	@Param			image			formData	file	true	"image to upload"
//	@Summary		Get All Images for a product
//	@Description	Retrieve All Images for the Product with This ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	string
//	@failure		default	{object}	[]db.ProductImage
//	@Router			/products/{id}/images [get]
func (h *ApiHandler) GetProductImages(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not parse product ID")
	}
	product, err := getProductHelper(c, int32(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not find product with ID: " + c.Params("id"))
	}
	images, _ := db.DB.GetProductImages(c.Context(), product.ID)
	return c.Status(fiber.StatusOK).JSON(images)
}

func getProductHelper(c *fiber.Ctx, productId int32) (db.Product, error) {
	product, err := db.DB.GetProduct(c.Context(), productId)
	if err != nil {
		return db.Product{}, err
	}
	return product, err
}

func populateProductParams(product ProductReq, params interface{}) error {
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
	case *db.CreateProductParams, *db.UpdateProductParams:
		var price, regPrice, disPrice *pgtype.Numeric
		var desc, shortDesc, pType, profileID *pgtype.Text
		var title *string
		var quantity *int32

		if createParams, ok := params.(*db.CreateProductParams); ok {
			price = &createParams.Price
			regPrice = &createParams.RegularPrice
			disPrice = &createParams.DiscountPrice
			desc = &createParams.Description
			shortDesc = &createParams.ShortDescription
			pType = &createParams.Type
			profileID = &createParams.ProfileID
			title = &createParams.Title
			quantity = &createParams.Quantity
		} else if updateParams, ok := params.(*db.UpdateProductParams); ok {
			price = &updateParams.Price
			regPrice = &updateParams.RegularPrice
			disPrice = &updateParams.DiscountPrice
			desc = &updateParams.Description
			shortDesc = &updateParams.ShortDescription
			pType = &updateParams.Type
			profileID = &updateParams.ProfileID
			title = &updateParams.Title
			quantity = &updateParams.Quantity
		}

		scanNumeric(price, product.Price)
		scanNumeric(regPrice, product.RegularPrice)
		scanNumeric(disPrice, product.DiscountPrice)
		scanText(desc, product.Description)
		scanText(shortDesc, product.ShortDescription)
		scanText(pType, product.Type)
		scanText(profileID, product.ProfileID)
		*title = product.Title
		*quantity = product.Quantity
		fmt.Println(p)
	default:
		return fmt.Errorf("unsupported params type")
	}

	if err != nil {
		return err
	}

	switch product.Type {
	case PRODUCT_RESERVATION, PRODUCT_DELISTED, PRODUCT_EXTERNAL, PRODUCT_BUY:
	default:
		return fmt.Errorf("product has unknown 'Type'")
	}

	return nil
}

func createProductHelper(product ProductReq) (db.CreateProductParams, error) {
	var params db.CreateProductParams
	err := populateProductParams(product, &params)
	return params, err
}

func updateProductHelper(product ProductReq) (db.UpdateProductParams, error) {
	var params db.UpdateProductParams
	err := populateProductParams(product, &params)
	return params, err
}
