CREATE TABLE characters_jutsus (
    id INTEGER PRIMARY KEY,
    character_id INTEGER NOT NULL,
    jutsu_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (character_id) REFERENCES characters(id) ON DELETE CASCADE,
    FOREIGN KEY (jutsu_id) REFERENCES jutsus(id) ON DELETE CASCADE,
    UNIQUE (character_id, jutsu_id)
);