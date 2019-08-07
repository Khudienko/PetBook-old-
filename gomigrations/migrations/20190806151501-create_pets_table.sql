-- +migrate Up
CREATE type sex AS ENUM('male','female');
CREATE type kind_of_animal AS ENUM('dog','cat','fish','bird','rodent');
CREATE table if not exists pets(
user_id integer not null primary key,
name text not null,
age integer not null,
animal_type kind_of_animal not null,
breed text not null,
description text ,
weight float not null,
gender sex not null
);
-- +migrate Down
DROP TABLE pets;
DROP type sex;
DROP type kind_of_animal;