CREATE DATABASE librady_db;

CREATE TABLE books(
    id bigserial,
    title varchar NOT NULL,
    description varchar NOT NULL,
    quantity int NOT NULL,
    cover varchar NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY (id)
);