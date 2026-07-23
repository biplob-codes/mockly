
-- name: CreateTeam :one
INSERT INTO teams (name)
VALUES (?)
RETURNING *;

-- name: ListTeams :many
SELECT * FROM teams
ORDER BY id;

-- name: GetTeam :one
SELECT * FROM teams
WHERE id = ?;

-- name: UpdateTeam :one
UPDATE teams
SET name = ?
WHERE id = ?
RETURNING *;

-- name: DeleteTeam :one
DELETE FROM teams
WHERE id = ?
RETURNING *;