CREATE DATABASE kernelpanicblog;

CREATE TABLE base_table (
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE users (
    id uuid PRIMARY KEY NOT NULL,
    username varchar(30) NOT NULL,
    hashedPassword varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    user_role varchar(15) NOT NULL
) INHERITS (base_table);

CREATE TABLE posts (
    post_id uuid PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    subtitle TEXT NOT NULL,
    content TEXT NOT NULL,
    slug TEXT NOT NULL,
    authorID uuid ,
    FOREIGN KEY (authorID) REFERENCES users (id)
) INHERITS (base_table);