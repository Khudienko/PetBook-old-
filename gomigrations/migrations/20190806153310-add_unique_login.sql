-- +migrate Up
ALTER TABLE users
    ADD constraint login_unique
    unique (login);
-- +migrate Down
ALTER TABLE users DROP CONSTRAINT login_unique;