
CREATE TABLE users (
    id uuid PRIMARY KEY NOT NULL,
    username varchar(30) NOT NULL,
    hashedPassword varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    user_role varchar(15) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE posts (
    post_id uuid PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    subtitle TEXT NOT NULL,
    content TEXT NOT NULL,
    slug TEXT NOT NULL,
    lang varchar(10) NOT NULL,
    authorID uuid NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (authorID) REFERENCES users (id)
);

CREATE TABLE comments (
    comment_id uuid PRIMARY KEY NOT NULL,
    comment_title VARCHAR(255) NOT NULL,
    comment_content TEXT NOT NULL,
    postID uuid NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (postID) REFERENCES posts (post_id)
);