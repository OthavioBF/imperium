-- name: CreateUser :one
INSERT INTO users ("user_name", "email", "password_hash", "bio")
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetUserById :one
SELECT 
    id,
    user_name,
    password_hash,
    email,
    bio,
    created_at,
    updated_at
FROM users
WHERE id = $1;


-- name: GetUsers :many
SELECT * FROM users;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET user_name = $1,
    email = $2,
    password_hash = $3,
    bio = $4
WHERE id = $5
RETURNING id;
