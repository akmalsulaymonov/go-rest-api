BEGIN;

CREATE TABLE IF NOT EXISTS comments (
    ID uuid PRIMARY KEY,
    Slug text NOT NULL,
    Author text NOT NULL,
    Body text NOT NULL
);

COMMIT;