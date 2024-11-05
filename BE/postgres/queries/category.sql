-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1;

-- name: ListCategories :many
SELECT *
FROM categories
ORDER BY name;

-- name: CreateCategory :one
INSERT INTO categories (name, description, parent_id)
VALUES ($1, $2, $3) RETURNING *;


-- name: GetCategoryBanner :one
SELECT banner_url FROM categories
WHERE id = $1;

-- name: GetCategoryBackground :one
SELECT background_url FROM categories
WHERE id = $1;

-- name: SetCategoryBanner :one
Update categories
SET banner_url = $1
WHERE id = $2 RETURNING *;

-- name: SetCategoryBackground :one
Update categories
SET background_url = $1
WHERE id = $2 RETURNING *;





-- name: UpdateCategory :one
UPDATE categories
SET name = $1, description = $2, parent_id = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $4 RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;



-- name: AddProductToCategory :one
INSERT INTO category_products (category_id, product_id)
VALUES ($1, $2) RETURNING *;

-- name: ListProductsInCategory :many
SELECT
    products.*
FROM
    categories
JOIN
    category_products ON categories.id = category_products.category_id
JOIN
    products ON products.id = category_products.product_id
WHERE 
    categories.id = $1;

-- name: ListAttributesInCategory :many
SELECT
    category_attributes.*
FROM
    categories
JOIN
    category_attributes ON categories.id = category_attributes.category_id
JOIN
    attribute_types ON category_attributes.attribute_type_id = attribute_types.id
WHERE 
    categories.id = $1;