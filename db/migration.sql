-- Création de la table User
CREATE TABLE IF NOT EXISTS User (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(255),
    email VARCHAR(255),
    `password` VARCHAR(255)
);

-- Création de la table Category
CREATE TABLE IF NOT EXISTS Category (
    category_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255)
);

-- Création de la table Post
CREATE TABLE IF NOT EXISTS Post (
    post_id INTEGER PRIMARY KEY AUTOINCREMENT,
    content TEXT,
    creation_date DATETIME,
    `user_id` INT,
    FOREIGN KEY (`user_id`) REFERENCES User(`user_id`)
);

-- Création de la table Comment
CREATE TABLE IF NOT EXISTS Comment (
    comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
    comment TEXT,
    date_creation DATETIME,
    post_id INT,
    `user_id` INT,
    FOREIGN KEY (post_id) REFERENCES Post(post_id),
    FOREIGN KEY (`user_id`) REFERENCES User(`user_id`)
);

-- Création de la table de liaison PostCategory
CREATE TABLE IF NOT EXISTS PostCategory (
    post_category_id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INT,
    category_id INT,
    FOREIGN KEY (post_id) REFERENCES Post(post_id),
    FOREIGN KEY (category_id) REFERENCES Category(category_id)
);

-- Création de la table de like - dislike
CREATE TABLE IF NOT EXISTS LikeDislike (
    like_dislike_id INTEGER PRIMARY KEY AUTOINCREMENT,
    `user_id` INTEGER,
    post_id INT,
    liked BOOLEAN,
    disliked BOOLEAN,
    FOREIGN KEY (`user_id`) REFERENCES User(`user_id`),
    FOREIGN KEY (post_id) REFERENCES Post(post_id)
);

/* CREATE TABLE IF NOT EXISTS Dislike (
    dislike_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INT,
    post_id INT,
    FOREIGN KEY (user_id) REFERENCES User(user_id),
    FOREIGN KEY (post_id) REFERENCES Post(post_id)
); */