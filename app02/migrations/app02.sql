-- +migrate Up
CREATE TABLE users(
    Id text primary key,
    Name varchar(20) ,
    Email text UNIQUE,
    Password text,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);
-- +migrate Down
DROP TABLE users;