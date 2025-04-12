CREATE TABLE IF NOT EXISTS users (
    id varchar PRIMARY KEY,
    first_name varchar NOT NULL,
    last_name varchar NOT NULL,
    email varchar NOT NULL,
    password varchar NOT NULL
)