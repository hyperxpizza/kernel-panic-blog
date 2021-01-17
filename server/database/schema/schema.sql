CREATE DATABASE kernelpanicblog;

CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL,
    username VARCHAR(30) NOT NULL,
    passwordHash TEXT NOT NULL,
    email TEXT NOT NULL,
    isAdmin BOOLEAN NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    description TEXT
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY NOT NULL,
    authorID INT,
    title TEXT NOT NULL,
    subtitle TEXT,
    content TEXT NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    slug TEXT NOT NULL,
    CONSTRAINT fk_authorID
        FOREIGN KEY (authorID)
        REFERENCES users(id)
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY NOT NULL,
    postID INT,
    content VARCHAR(1000),
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    isAdmin BOOLEAN NOT NULL,
    opEmail VARCHAR(200) NOT NULL,
    opName VARCHAR(200) NOT NULL,
    CONSTRAINT fk_postID
        FOREIGN KEY (postID)
        REFERENCES posts(id)
);

CREATE TABLE tagmap (
    id SERIAL PRIMARY KEY NOT NULL,
    postID INT, 
    tagID INT,
    CONSTRAINT fk_postID
    FOREIGN KEY (postID)
    REFERENCES posts(id),
    CONSTRAINT fk_tagID
    FOREIGN KEY(tagID)
    REFERENCES tags(id) 
);

CREATE TABLE tags (
    id SERIAL PRIMARY KEY NOT NULL,
    tagName VARCHAR(100) NOT NULL,
    slug TEXT NOT NULL
);