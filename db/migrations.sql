-- name: users-table
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255),
    date_created DATE,
    is_admin BOOLEAN
);

-- name: user-auth-table
CREATE TABLE user_auth (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    hash VARCHAR(255),
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- name: posts-table
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    body TEXT,
    thumbnail_url VARCHAR(255),
    creator_id INT NOT NULL,
    time_created DATETIME NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES users (id)
);

-- name: comments-table
CREATE TABLE comments (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    body TEXT,
    creator_id INT NOT NULL,
    post_id INT NOT NULL,
    time_created DATETIME NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES users (id),
    FOREIGN KEY (post_id) REFERENCES posts (id)
);

-- name: likes-table
CREATE TABLE likes (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    creator_id INT NOT NULL,
    post_id INT NOT NULL,
    time_created DATETIME NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES users (id),
    FOREIGN KEY (post_id) REFERENCES posts (id)
);
