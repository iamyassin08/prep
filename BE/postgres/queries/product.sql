-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- -- name: GetUserCategory :many
-- SELECT sqlc.embed(users), sqlc.embed(categories)
-- FROM users
-- JOIN categories ON users.category_id = categories.id
-- WHERE users.id = $1;

-- name: CreateUserImage :one
INSERT INTO user_images (user_id, image_url)
VALUES ($1, $2) RETURNING *;

-- name: GetUserImage :one
SELECT * FROM user_images
WHERE user_id = $1 AND image_url = $2;

-- name: GetUserCategory :one
SELECT  
    categories.*
FROM  
    category_users
JOIN 
    categories ON category_users.category_id  = categories.id
WHERE 
    category_users.user_id = $1;



-- name: GetUserImages :many
SELECT * FROM user_images
WHERE user_id = $1;

-- name: SetUserThumbnail :one
INSERT INTO user_image_thumbnail (user_id, user_image_id)
VALUES ($1, $2) RETURNING *;

-- name: UpdateUserThumbnail :exec
UPDATE user_image_thumbnail
SET user_image_id = $1
WHERE user_id = $2;

-- name: GetUserThumbnail :one
SELECT * FROM user_image_thumbnail
WHERE user_id = $1;

-- name: GetUserThumbnailImage :one
SELECT * FROM user_images
WHERE id = $1;


-- name: CreateUser :one
INSERT INTO users (title, description, short_description, price, quantity, regular_price, discount_price, profile_id, type)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET title = $1, description = $2, short_description = $3, price = $4, quantity = $5, regular_price = $6, discount_price = $7, profile_id = $8, type = $9, updated_at = CURRENT_TIMESTAMP
WHERE id = $10;


-- name: UpdateUserDescription :exec
UPDATE users
SET description = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateUserQuantity :exec
UPDATE users
SET quantity = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateUserFirstName :exec
UPDATE users
SET title = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateUserEmail  :exec
UPDATE users
SET price = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateUserLastName  :exec
UPDATE users
SET regular_price = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateUserDiscountEmail  :exec
UPDATE users
SET discount_price = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY title;









