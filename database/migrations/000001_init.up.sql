CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS leak_data
(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    status DATA_STATUS,
    data JSONB

);

