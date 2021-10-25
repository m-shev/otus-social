CREATE TABLE user
(
    id       INT AUTO_INCREMENT PRIMARY KEY,
    name     VARCHAR(255) NOT NULL,
    surname  VARCHAR(255) NOT NULL,
    avatar   VARCHAR(255) DEFAULT '',
    age      TINYINT UNSIGNED NOT NULL,
    gender   ENUM('male', 'female') NOT NULL,
    city     VARCHAR(255) NOT NULL,
    email    VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);