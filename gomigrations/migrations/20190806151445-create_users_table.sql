-- +migrate Up
CREATE type role AS ENUM('pet','vet');
CREATE TABLE IF NOT EXISTS users(
id serial not null,
email text not null,
firstname text not null,
lastname text not null ,
login text not null ,
pet_or_vet role,
password text not null
);
-- +migrate Down
DROP TABLE users;
DROP type role;