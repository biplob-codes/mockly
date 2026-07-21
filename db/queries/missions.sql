-- name: CreateMission :one
INSERT INTO missions (name, description, assigned_to, rank, status, reward, starts_at)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetMission :one
SELECT * FROM missions
WHERE id = ?;

-- name: ListMissions :many
SELECT * FROM missions;

-- name: GetMissionsByTeam :many
SELECT * FROM missions
WHERE assigned_to = ?;

-- name: DeleteMission :one
DELETE FROM missions
WHERE id = ?
RETURNING *;