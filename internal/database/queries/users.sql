-- name: GetUser :one
SELECT * FROM users WHERE id = ?  LIMIT 1;

-- name: CreateUser :one
INSERT INTO users(username,full_name,email)
VALUES (?,?,?)
RETURNING * ;