CREATE TABLE jutsus (
    id          INTEGER PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT,
    type        TEXT NOT NULL CHECK (type IN ('Ninjutsu','Taijutsu','Genjutsu','Kekkei Genkai','Senjutsu')),
    rank        TEXT NOT NULL CHECK (rank IN ('E','D','C','B','A','S')),
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);