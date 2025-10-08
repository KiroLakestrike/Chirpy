-- name: UpdateUser :one
UPDATE users
SET email = $1, hashed_passwords = $2, updated_at = NOW()
WHERE id = $3
RETURNING *;