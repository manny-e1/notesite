CREATE TABLE users (
    id  integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(255) NOT NULL,
    username varchar(255) UNIQUE NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    gender varchar(255) NOT NULL,
    role varchar(255) NOT NULL,
    course varchar(255) NOT NULL,
    image varchar(225) DEFAULT 'profile.jpg',
    about varchar(300) NOT NULL DEFAULT 'N/A',
    joindate timestamp NOT NULL
); 


CREATE TABLE posts(
id  integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
user_id integer REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
text  varchar(300),
posted_at timestamp ,
updated_at timestamp
);
CREATE TABLE comments
(
    id         integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    writer     integer REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    comment    varchar(300) NOT NULL,
    post_id    integer REFERENCES posts(id),
    commented_at    TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
CREATE TABLE uploads(
id  integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
title varchar(255) NOT NULL,
file_description text NOT NULL,
file_type varchar(255) NOT NULL,
file_uploader varchar(255) REFERENCES users(username) ON UPDATE CASCADE ON DELETE CASCADE,
file_uploader_ID integer REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
file_uploaded_on timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
file_uploaded_to varchar(255) NOT NULL,
file_path varchar(225) NOT NULL,
file_status varchar(255) NOT NULL DEFAULT 'not approved yet'
);

