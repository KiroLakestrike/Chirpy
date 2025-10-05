-- name: GetUserByEMail :one
SELECT *
FROM users
WHERE email = $1;