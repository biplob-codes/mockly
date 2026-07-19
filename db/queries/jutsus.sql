-- name: CreateJutsu :one
INSERT INTO jutsus (name, description, type, rank)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetJutsus :many
SELECT * FROM jutsus;

-- name: GetJutsu :one
SELECT * FROM jutsus WHERE id = ?;

-- name: DeleteJutsu :one
DELETE FROM jutsus WHERE id = ? RETURNING *;