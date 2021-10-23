create table user_friend (
    userFrom INT NOT NULL,
    userTo INT NOT NULL,
    reference ENUM('friend', 'response', 'request', 'refusal'),
    FOREIGN key (userFrom) REFERENCES user (id),
    FOREIGN key (userTo) REFERENCES user (id),
    PRIMARY KEY (userFrom, userTo)
)