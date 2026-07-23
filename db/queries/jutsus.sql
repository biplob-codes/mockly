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

-- name: UpdateJutsu :one
UPDATE jutsus
SET
  name        = COALESCE(sqlc.narg('name'), name),
  description = COALESCE(sqlc.narg('description'), description),
  type        = COALESCE(sqlc.narg('type'), type),
  rank        = COALESCE(sqlc.narg('rank'), rank),
  updated_at  = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('id')
RETURNING *;