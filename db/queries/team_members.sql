-- name: GetTeamDetails :many
SELECT
    t.id AS team_id,
    t.name AS team_name,
    t.created_at AS team_created_at,
    c.id, c.name, c.nickname, c.clan, c.age, c.rank, c.birthdate, c.village_id, c.created_at, c.updated_at,
    tm.role,
    tm.joined_at
FROM teams t
LEFT JOIN team_members tm ON tm.team_id = t.id
LEFT JOIN characters c ON c.id = tm.character_id
WHERE t.id = ?;

-- name: GetCharacterMembership :one
SELECT * FROM team_members
WHERE character_id = ?;

-- name: GetTeamSensei :one
SELECT * FROM team_members
WHERE team_id = ? AND role = 'sensei';

-- name: AddTeamMember :one
INSERT INTO team_members (team_id, character_id, role)
VALUES (?, ?, ?)
RETURNING *;

-- name: RemoveTeamMember :one
DELETE FROM team_members
WHERE team_id = ? AND character_id = ?
RETURNING *;

-- name: CharacterTeamCount :one
SELECT COUNT(*) FROM team_members WHERE character_id = ?;