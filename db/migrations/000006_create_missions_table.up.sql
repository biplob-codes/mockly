CREATE TABLE missions (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    assigned_to INTEGER NOT NULL,
    rank TEXT NOT NULL CHECK (rank IN ('D', 'C', 'B', 'A', 'S')),
    status TEXT NOT NULL DEFAULT 'Pending' CHECK (status IN ('Pending', 'Failed', 'Success')),
    reward INTEGER NOT NULL,
    starts_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (assigned_to) REFERENCES teams(id) ON DELETE CASCADE
);