-- name: GetUserByID :one
SELECT *
FROM users
WHERE ID = $1;