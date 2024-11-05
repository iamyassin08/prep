// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: product.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (title, description, short_description, price, quantity, regular_price, discount_price, profile_id, type)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, profile_id, title, description, short_description, price, quantity, discount_price, regular_price, created_at, updated_at, type
`

type CreateProductParams struct {
	Title            string
	Description      pgtype.Text
	ShortDescription pgtype.Text
	Price            pgtype.Numeric
	Quantity         int32
	RegularPrice     pgtype.Numeric
	DiscountPrice    pgtype.Numeric
	ProfileID        pgtype.Text
	Type             pgtype.Text
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.Title,
		arg.Description,
		arg.ShortDescription,
		arg.Price,
		arg.Quantity,
		arg.RegularPrice,
		arg.DiscountPrice,
		arg.ProfileID,
		arg.Type,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.ProfileID,
		&i.Title,
		&i.Description,
		&i.ShortDescription,
		&i.Price,
		&i.Quantity,
		&i.DiscountPrice,
		&i.RegularPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
	)
	return i, err
}

const createProductImage = `-- name: CreateProductImage :one

INSERT INTO product_images (product_id, image_url)
VALUES ($1, $2) RETURNING id, product_id, image_url, created_at
`

type CreateProductImageParams struct {
	ProductID int32
	ImageUrl  string
}

// -- name: GetProductCategory :many
// SELECT sqlc.embed(products), sqlc.embed(categories)
// FROM products
// JOIN categories ON products.category_id = categories.id
// WHERE products.id = $1;
func (q *Queries) CreateProductImage(ctx context.Context, arg CreateProductImageParams) (ProductImage, error) {
	row := q.db.QueryRow(ctx, createProductImage, arg.ProductID, arg.ImageUrl)
	var i ProductImage
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, profile_id, title, description, short_description, price, quantity, discount_price, regular_price, created_at, updated_at, type FROM products
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.ProfileID,
		&i.Title,
		&i.Description,
		&i.ShortDescription,
		&i.Price,
		&i.Quantity,
		&i.DiscountPrice,
		&i.RegularPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
	)
	return i, err
}

const getProductCategory = `-- name: GetProductCategory :one
SELECT  
    categories.id, categories.parent_id, categories.name, categories.description, categories.banner_url, categories.background_url, categories.updated_at, categories.created_at
FROM  
    category_products
JOIN 
    categories ON category_products.category_id  = categories.id
WHERE 
    category_products.product_id = $1
`

func (q *Queries) GetProductCategory(ctx context.Context, productID int32) (Category, error) {
	row := q.db.QueryRow(ctx, getProductCategory, productID)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.ParentID,
		&i.Name,
		&i.Description,
		&i.BannerUrl,
		&i.BackgroundUrl,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getProductImage = `-- name: GetProductImage :one
SELECT id, product_id, image_url, created_at FROM product_images
WHERE product_id = $1 AND image_url = $2
`

type GetProductImageParams struct {
	ProductID int32
	ImageUrl  string
}

func (q *Queries) GetProductImage(ctx context.Context, arg GetProductImageParams) (ProductImage, error) {
	row := q.db.QueryRow(ctx, getProductImage, arg.ProductID, arg.ImageUrl)
	var i ProductImage
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getProductImages = `-- name: GetProductImages :many
SELECT id, product_id, image_url, created_at FROM product_images
WHERE product_id = $1
`

func (q *Queries) GetProductImages(ctx context.Context, productID int32) ([]ProductImage, error) {
	rows, err := q.db.Query(ctx, getProductImages, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductImage
	for rows.Next() {
		var i ProductImage
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.ImageUrl,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductThumbnail = `-- name: GetProductThumbnail :one
SELECT product_id, product_image_id, created_at FROM product_image_thumbnail
WHERE product_id = $1
`

func (q *Queries) GetProductThumbnail(ctx context.Context, productID int32) (ProductImageThumbnail, error) {
	row := q.db.QueryRow(ctx, getProductThumbnail, productID)
	var i ProductImageThumbnail
	err := row.Scan(&i.ProductID, &i.ProductImageID, &i.CreatedAt)
	return i, err
}

const getProductThumbnailImage = `-- name: GetProductThumbnailImage :one
SELECT id, product_id, image_url, created_at FROM product_images
WHERE id = $1
`

func (q *Queries) GetProductThumbnailImage(ctx context.Context, id int32) (ProductImage, error) {
	row := q.db.QueryRow(ctx, getProductThumbnailImage, id)
	var i ProductImage
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, profile_id, title, description, short_description, price, quantity, discount_price, regular_price, created_at, updated_at, type FROM products
ORDER BY title
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.ProfileID,
			&i.Title,
			&i.Description,
			&i.ShortDescription,
			&i.Price,
			&i.Quantity,
			&i.DiscountPrice,
			&i.RegularPrice,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Type,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setProductThumbnail = `-- name: SetProductThumbnail :one
INSERT INTO product_image_thumbnail (product_id, product_image_id)
VALUES ($1, $2) RETURNING product_id, product_image_id, created_at
`

type SetProductThumbnailParams struct {
	ProductID      int32
	ProductImageID int32
}

func (q *Queries) SetProductThumbnail(ctx context.Context, arg SetProductThumbnailParams) (ProductImageThumbnail, error) {
	row := q.db.QueryRow(ctx, setProductThumbnail, arg.ProductID, arg.ProductImageID)
	var i ProductImageThumbnail
	err := row.Scan(&i.ProductID, &i.ProductImageID, &i.CreatedAt)
	return i, err
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET title = $1, description = $2, short_description = $3, price = $4, quantity = $5, regular_price = $6, discount_price = $7, profile_id = $8, type = $9, updated_at = CURRENT_TIMESTAMP
WHERE id = $10
`

type UpdateProductParams struct {
	Title            string
	Description      pgtype.Text
	ShortDescription pgtype.Text
	Price            pgtype.Numeric
	Quantity         int32
	RegularPrice     pgtype.Numeric
	DiscountPrice    pgtype.Numeric
	ProfileID        pgtype.Text
	Type             pgtype.Text
	ID               int32
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.Exec(ctx, updateProduct,
		arg.Title,
		arg.Description,
		arg.ShortDescription,
		arg.Price,
		arg.Quantity,
		arg.RegularPrice,
		arg.DiscountPrice,
		arg.ProfileID,
		arg.Type,
		arg.ID,
	)
	return err
}

const updateProductDescription = `-- name: UpdateProductDescription :exec
UPDATE products
SET description = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2
`

type UpdateProductDescriptionParams struct {
	Description pgtype.Text
	ID          int32
}

func (q *Queries) UpdateProductDescription(ctx context.Context, arg UpdateProductDescriptionParams) error {
	_, err := q.db.Exec(ctx, updateProductDescription, arg.Description, arg.ID)
	return err
}

const updateProductDiscountPrice = `-- name: UpdateProductDiscountPrice :exec
UPDATE products
SET discount_price = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2
`

type UpdateProductDiscountPriceParams struct {
	DiscountPrice pgtype.Numeric
	ID            int32
}

func (q *Queries) UpdateProductDiscountPrice(ctx context.Context, arg UpdateProductDiscountPriceParams) error {
	_, err := q.db.Exec(ctx, updateProductDiscountPrice, arg.DiscountPrice, arg.ID)
	return err
}

const updateProductName = `-- name: UpdateProductName :exec
UPDATE products
SET title = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2
`

type UpdateProductNameParams struct {
	Title string
	ID    int32
}

func (q *Queries) UpdateProductName(ctx context.Context, arg UpdateProductNameParams) error {
	_, err := q.db.Exec(ctx, updateProductName, arg.Title, arg.ID)
	return err
}

const updateProductPrice = `-- name: UpdateProductPrice :exec
UPDATE products
SET price = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2
`

type UpdateProductPriceParams struct {
	Price pgtype.Numeric
	ID    int32
}

func (q *Queries) UpdateProductPrice(ctx context.Context, arg UpdateProductPriceParams) error {
	_, err := q.db.Exec(ctx, updateProductPrice, arg.Price, arg.ID)
	return err
}

const updateProductQuantity = `-- name: UpdateProductQuantity :exec
UPDATE products
SET quantity = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2
`

type UpdateProductQuantityParams struct {
	Quantity int32
	ID       int32
}

func (q *Queries) UpdateProductQuantity(ctx context.Context, arg UpdateProductQuantityParams) error {
	_, err := q.db.Exec(ctx, updateProductQuantity, arg.Quantity, arg.ID)
	return err
}

const updateProductRegularPrice = `-- name: UpdateProductRegularPrice :exec
UPDATE products
SET regular_price = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2
`

type UpdateProductRegularPriceParams struct {
	RegularPrice pgtype.Numeric
	ID           int32
}

func (q *Queries) UpdateProductRegularPrice(ctx context.Context, arg UpdateProductRegularPriceParams) error {
	_, err := q.db.Exec(ctx, updateProductRegularPrice, arg.RegularPrice, arg.ID)
	return err
}

const updateProductThumbnail = `-- name: UpdateProductThumbnail :exec
UPDATE product_image_thumbnail
SET product_image_id = $1
WHERE product_id = $2
`

type UpdateProductThumbnailParams struct {
	ProductImageID int32
	ProductID      int32
}

func (q *Queries) UpdateProductThumbnail(ctx context.Context, arg UpdateProductThumbnailParams) error {
	_, err := q.db.Exec(ctx, updateProductThumbnail, arg.ProductImageID, arg.ProductID)
	return err
}
