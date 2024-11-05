-- name: AddUserProfile :exec
INSERT INTO profiles (id)
VALUES ($1);
-- name: GetUserProfile :one
SELECT * FROM profiles
WHERE id = $1 LIMIT 1;

-- name: AddUserFavorite :exec
INSERT INTO profile_user_favorites (user_id, profile_id)
VALUES ($1, $2);
-- name: RemoveUserfavorite :exec
DELETE FROM profile_user_favorites
WHERE user_id = $1 AND profile_id = $2;

-- -- name: GetUserFavorites :many
-- SELECT * FROM profile_user_favorites
-- WHERE profile_id = $1;

-- name: GetUserFavorites :many
SELECT
    users.*
FROM
    users
JOIN
    profile_user_favorites ON profile_user_favorites.user_id = users.id
WHERE 
    profile_user_favorites.profile_id = $1;





-- name: AddUserCart :exec
INSERT INTO profile_user_cart_items (user_id, profile_id, user_quantity)
VALUES ($1, $2, $3);
-- name: RemoveUserCart :exec
DELETE FROM profile_user_cart_items
WHERE user_id = $1 AND profile_id = $2;
-- name: UpdateUserCartQuantity :exec
UPDATE profile_user_cart_items
SET user_quantity = $1
WHERE user_id = $2 AND profile_id = $3;

-- name: GetUserInUserCart :one
SELECT * FROM profile_user_cart_items
WHERE user_id = $1 AND profile_id = $2;

-- name: GetUserCartItems :many
SELECT
    users.id,
    users.profile_id,
    users.title,
    users.short_description,
    users.price,
    profile_user_cart_items.user_quantity AS quantity
FROM
    users
JOIN
    profile_user_cart_items ON profile_user_cart_items.user_id = users.id
WHERE 
    profile_user_cart_items.profile_id = $1;


