CREATE TABLE subscribers 
(
    id serial  unique not null,
    email  varchar(255) unique not null,
    password_hash varchar(255) not null
);