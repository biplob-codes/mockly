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