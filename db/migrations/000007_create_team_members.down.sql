PRAGMA foreign_keys=off;

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
    team_id    INTEGER REFERENCES teams(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (village_id) REFERENCES villages(id) ON DELETE RESTRICT
);

DROP TABLE teams;
CREATE TABLE teams (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    sensei_id INTEGER NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sensei_id) REFERENCES characters(id) ON DELETE RESTRICT
);

DROP TABLE team_members;

PRAGMA foreign_keys=on;