CREATE TABLE IF NOT EXISTS files (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    size BIGINT NOT NULL,
    path text NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_files_name ON files(name);