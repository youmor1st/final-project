CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE,
    password VARCHAR(50),
    role VARCHAR(20)
    );

CREATE TABLE IF NOT EXISTS merits (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    points INT,
    reason TEXT,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
