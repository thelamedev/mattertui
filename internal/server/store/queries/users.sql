-- name: CreateUser :one
INSERT INTO users (
  username, email, password_hash
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserPasswordVersionByUserID :one
SELECT password_version FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;
