CREATE TABLE dialog_member
(
    member_id INT UNIQUE,
    dialog_id BIGINT,
    role ENUM('creator', 'member') NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (dialog_id) REFERENCES dialog (dialog_id),
    PRIMARY KEY (dialog_id, member_id)
)