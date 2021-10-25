create table user_interest (
  userId INT,
  interestId INT,
  FOREIGN KEY (userId) REFERENCES user (id),
  FOREIGN KEY (interestId) REFERENCES interest (id) ON DELETE CASCADE,
  PRIMARY KEY (userId, interestId)
);