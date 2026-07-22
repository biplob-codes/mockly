-- name: GetVillages :many
SELECT * FROM villages;

-- name: GetVillage :one
SELECT * FROM villages WHERE id=?;

-- name: CreateVillage :one
INSERT INTO villages (name,description,land,population,kage_id,founded_at)
VALUES (?,?,?,?,?,?) 
RETURNING *;

-- name: DeleteVillage :one
DELETE FROM villages WHERE id=? RETURNING *;

-- name: UpdateVillage :one
UPDATE villages
SET
  name        = COALESCE(sqlc.narg('name'), name),
  description = COALESCE(sqlc.narg('description'), description),
  land        = COALESCE(sqlc.narg('land'), land),
  population  = COALESCE(sqlc.narg('population'), population),
  kage_id     = COALESCE(sqlc.narg('kage_id'), kage_id),
  founded_at  = COALESCE(sqlc.narg('founded_at'), founded_at),
  updated_at  = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('id')
RETURNING *;