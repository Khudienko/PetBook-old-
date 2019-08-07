-- +migrate Up
ALTER TABLE users
ADD CONSTRAINT PK_ID PRIMARY KEY (id);
-- +migrate Down
ALTER TABLE users DROP CONSTRAINT PK_ID;
