-- name: CreateTeam :one
INSERT INTO teams (name, sensei_id)
VALUES (?, ?)
RETURNING *;

-- name: DeleteTeam :one
DELETE FROM teams
WHERE id = ?
RETURNING *;

-- name: GetTeam :one
SELECT * FROM teams
WHERE id = ?;

-- name: GetTeamMembers :many
SELECT * FROM characters
WHERE team_id = ?;

-- name: ListTeams :many
SELECT * FROM teams; 