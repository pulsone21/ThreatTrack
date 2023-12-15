CREATE TABLE IF NOT EXISTS    task_comments (
        id varchar(255) NOT NULL COMMENT 'UUID',
        create_time timestamp NULL DEFAULT NULL COMMENT 'Create Time',
        content text NOT NULL,
        writer varchar(255) NOT NULL,
        task varchar(255) NOT NULL,
        PRIMARY KEY (id),
        KEY writer (writer),
        KEY task (task),
        FOREIGN KEY (writer) REFERENCES users (id),
        FOREIGN KEY (task) REFERENCES task (id)
    ) 