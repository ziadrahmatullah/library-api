CREATE DATABASE library_db;

CREATE TABLE authors(
    id bigserial,
    name varchar NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE users(
    id bigserial,
    name varchar NOT NULL,
    email varchar NOT NULL,
    phone varchar NOT NULL,
    password varchar NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE books(
    id bigserial,
    title varchar NOT NULL,
    description varchar NOT NULL,
    quantity int NOT NULL,
    cover varchar,
    author_id bigint NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY (id),
    FOREIGN KEY (author_id) REFERENCES authors(id)
);

CREATE TABLE borrowing_books(
    id bigserial,
    book_id bigint NOT NULL,
    user_id bigint NOT NULL,
    status varchar NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY (id),
    FOREIGN KEY (book_id) REFERENCES books(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

