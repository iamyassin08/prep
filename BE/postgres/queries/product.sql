-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- -- name: GetProductCategory :many
-- SELECT sqlc.embed(products), sqlc.embed(categories)
-- FROM products
-- JOIN categories ON products.category_id = categories.id
-- WHERE products.id = $1;

-- name: CreateProductImage :one
INSERT INTO product_images (product_id, image_url)
VALUES ($1, $2) RETURNING *;

-- name: GetProductImage :one
SELECT * FROM product_images
WHERE product_id = $1 AND image_url = $2;

-- name: GetProductCategory :one
SELECT  
    categories.*
FROM  
    category_products
JOIN 
    categories ON category_products.category_id  = categories.id
WHERE 
    category_products.product_id = $1;



-- name: GetProductImages :many
SELECT * FROM product_images
WHERE product_id = $1;

-- name: SetProductThumbnail :one
INSERT INTO product_image_thumbnail (product_id, product_image_id)
VALUES ($1, $2) RETURNING *;

-- name: UpdateProductThumbnail :exec
UPDATE product_image_thumbnail
SET product_image_id = $1
WHERE product_id = $2;

-- name: GetProductThumbnail :one
SELECT * FROM product_image_thumbnail
WHERE product_id = $1;

-- name: GetProductThumbnailImage :one
SELECT * FROM product_images
WHERE id = $1;


-- name: CreateProduct :one
INSERT INTO products (title, description, short_description, price, quantity, regular_price, discount_price, profile_id, type)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;

-- name: UpdateProduct :exec
UPDATE products
SET title = $1, description = $2, short_description = $3, price = $4, quantity = $5, regular_price = $6, discount_price = $7, profile_id = $8, type = $9, updated_at = CURRENT_TIMESTAMP
WHERE id = $10;


-- name: UpdateProductDescription :exec
UPDATE products
SET description = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateProductQuantity :exec
UPDATE products
SET quantity = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateProductName :exec
UPDATE products
SET title = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateProductPrice :exec
UPDATE products
SET price = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateProductRegularPrice :exec
UPDATE products
SET regular_price = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateProductDiscountPrice :exec
UPDATE products
SET discount_price = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY title;









