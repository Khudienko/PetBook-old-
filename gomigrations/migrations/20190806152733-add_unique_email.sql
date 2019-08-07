-- +migrate Up
ALTER TABLE users
    Add constraint email_unique
    unique (email);
-- +migrate Down
ALTER TABLE users DROP CONSTRAINT email_unique;