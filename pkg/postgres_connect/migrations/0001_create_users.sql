-- +migrate Up
CREATE TABLE users(
    user_id SERIAL PRIMARY KEY ,
    login VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
);

-- +migrate Down
DROP TABLE users;