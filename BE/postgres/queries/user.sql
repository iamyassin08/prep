-- name: AddUserProfile :exec
INSERT INTO profiles (id)
VALUES ($1);
-- name: GetUserProfile :one
SELECT * FROM profiles
WHERE id = $1 LIMIT 1;

-- name: AddProductFavorite :exec
INSERT INTO profile_product_favorites (product_id, profile_id)
VALUES ($1, $2);
-- name: RemoveProductfavorite :exec
DELETE FROM profile_product_favorites
WHERE product_id = $1 AND profile_id = $2;

-- -- name: GetUserFavorites :many
-- SELECT * FROM profile_product_favorites
-- WHERE profile_id = $1;

-- name: GetUserFavorites :many
SELECT
    products.*
FROM
    products
JOIN
    profile_product_favorites ON profile_product_favorites.product_id = products.id
WHERE 
    profile_product_favorites.profile_id = $1;





-- name: AddProductCart :exec
INSERT INTO profile_product_cart_items (product_id, profile_id, product_quantity)
VALUES ($1, $2, $3);
-- name: RemoveProductCart :exec
DELETE FROM profile_product_cart_items
WHERE product_id = $1 AND profile_id = $2;
-- name: UpdateProductCartQuantity :exec
UPDATE profile_product_cart_items
SET product_quantity = $1
WHERE product_id = $2 AND profile_id = $3;

-- name: GetProductInUserCart :one
SELECT * FROM profile_product_cart_items
WHERE product_id = $1 AND profile_id = $2;

-- name: GetUserCartItems :many
SELECT
    products.id,
    products.profile_id,
    products.title,
    products.short_description,
    products.price,
    profile_product_cart_items.product_quantity AS quantity
FROM
    products
JOIN
    profile_product_cart_items ON profile_product_cart_items.product_id = products.id
WHERE 
    profile_product_cart_items.profile_id = $1;


