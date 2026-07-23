CREATE TABLE team_members (
    team_id INTEGER NOT NULL,
    character_id INTEGER NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('Sensei', 'Student')),
    joined_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (team_id, character_id),
    FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE,
    FOREIGN KEY (character_id) REFERENCES characters(id) ON DELETE CASCADE
);
ALTER TABLE teams DROP COLUMN sensei_id;
ALTER TABLE characters DROP COLUMN  team_id;