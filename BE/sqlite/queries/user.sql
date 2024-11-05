-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email)
VALUES (?, ?, ?) RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET first_name = ?, last_name = ?, email = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY first_name;
