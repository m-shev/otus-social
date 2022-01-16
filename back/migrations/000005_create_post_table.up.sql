create table post (
    id int AUTO_INCREMENT PRIMARY KEY,
    authorId int NOT NULL,
    content text NOT NULL,
    imageLink varchar(4096),
    createAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updateAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (authorId) REFERENCES user (id)
)