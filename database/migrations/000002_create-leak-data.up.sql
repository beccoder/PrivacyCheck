DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'data_status') THEN
            CREATE TYPE DATA_STATUS AS ENUM ('FOUND', 'NOT_FOUND');
        END IF;
    END $$;

CREATE TABLE IF NOT EXISTS leak_data
(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    status DATA_STATUS NOT NULL,
    data JSONB
);
