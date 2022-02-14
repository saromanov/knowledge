CREATE TABLE IF NOT EXISTS article(
    id serial PRIMARY KEY;
    created_at timestamp;
    updated_at timestamp;
    title text NOT NULL;
    body text;
    author_id text NOT NULL;
);

CREATE TABLE IF NOT EXISTS author(
    id serial PRIMARY KEY;
    name text;
);

CREATE TABLE IF NOT EXISTS link(
    id serial PRIMARY KEY;
    url text NOT NULL;
);