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