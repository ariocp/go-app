CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    email VARCHAR(255),
    confirmation_code VARCHAR(4),
    is_confirmed BOOLEAN DEFAULT FALSE,
    confirmation_expiry TIMESTAMP
);

CREATE UNIQUE INDEX idx_username ON users(username);