-- name: CreateCharacterJutsu :one
INSERT INTO characters_jutsus (character_id, jutsu_id)
VALUES (?, ?)
RETURNING *;

-- name: DeleteCharacterJutsu :exec
DELETE FROM characters_jutsus
WHERE character_id = ? AND jutsu_id = ?
RETURNING *;

-- name: GetJutsusByCharacterID :many
SELECT jutsus.*
FROM jutsus
JOIN characters_jutsus ON characters_jutsus.jutsu_id = jutsus.id
WHERE characters_jutsus.character_id = ?;

-- name: GetCharactersByJutsuID :many
SELECT characters.*
FROM characters
JOIN characters_jutsus ON characters_jutsus.character_id = characters.id
WHERE characters_jutsus.jutsu_id = ?;