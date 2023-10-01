CREATE TABLE if NOT EXISTS albums (
    id serial NOT NULL PRIMARY KEY,
    album json NOT NULL
);