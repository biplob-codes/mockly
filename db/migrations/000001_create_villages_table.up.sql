CREATE TABLE villages (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  description TEXT,
  land TEXT NOT NULL CHECK (land IN ('Fire','Wind','Water','Earth','Lightning')),
  population INTEGER,
  kage_id INTEGER REFERENCES characters(id),
  founded_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);