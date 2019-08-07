
-- +migrate Up
ALTER TABLE pets
    ADD constraint pets_fk
    FOREIGN KEY (user_id) references users(id);
-- +migrate Down
ALTER TABLE pets DROP CONSTRAINT pets_fk;