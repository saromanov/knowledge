BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS author(
    ID INT PRIMARY KEY     NOT NULL,
    name text not null,
    created_at date not null
);

CREATE TABLE IF NOT EXISTS "page" (
    ID INT PRIMARY KEY     NOT NULL,
    created_at date not null,
    updated_at date not null,
    title text,
    body text,
    author_id int
)

CREATE TABLE IF NOT EXISTS link(
    id serial PRIMARY KEY,
    url text NOT NULL
);
COMMIT;