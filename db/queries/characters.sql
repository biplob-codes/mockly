-- name: GetCharacters :many
SELECT * FROM characters;

-- name: GetCharacter :one
SELECT * FROM characters WHERE id = ?;

-- name: CreateCharacter :one
INSERT INTO characters (name, nickname, clan, age, rank, birthdate, village_id)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteCharacter :one
DELETE FROM characters WHERE id = ? RETURNING *;

-- name: UpdateCharacter :one
UPDATE characters
SET
  name       = COALESCE(sqlc.narg('name'), name),
  nickname   = COALESCE(sqlc.narg('nickname'), nickname),
  clan       = COALESCE(sqlc.narg('clan'), clan),
  age        = COALESCE(sqlc.narg('age'), age),
  rank       = COALESCE(sqlc.narg('rank'), rank),
  birthdate  = COALESCE(sqlc.narg('birthdate'), birthdate),
  village_id = COALESCE(sqlc.narg('village_id'), village_id),
  updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('id')
RETURNING *;