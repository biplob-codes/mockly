PRAGMA foreign_keys=off;

CREATE TABLE team_members (
    team_id INTEGER NOT NULL,
    character_id INTEGER NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('Sensei', 'Student')),
    joined_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (team_id, character_id),
    FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE,
    FOREIGN KEY (character_id) REFERENCES characters(id) ON DELETE CASCADE
);

DROP TABLE teams;
CREATE TABLE teams (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE characters;
CREATE TABLE characters (
    id         INTEGER PRIMARY KEY,
    name       TEXT NOT NULL,
    nickname   TEXT,
    clan       TEXT,
    age        INTEGER,
    rank       TEXT NOT NULL CHECK (rank IN ('Genin','Chunin','Jonin','Kage')),
    birthdate  DATE NOT NULL,
    village_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (village_id) REFERENCES villages(id) ON DELETE RESTRICT
);

PRAGMA foreign_keys=on;