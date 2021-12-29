CREATE TABLE message
(
    message_id BINARY(16) PRIMARY KEY,
    dialog_id  BIGINT NOT NULL,
    author_id INT NOT NULL,
    content TEXT,
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)