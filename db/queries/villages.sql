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