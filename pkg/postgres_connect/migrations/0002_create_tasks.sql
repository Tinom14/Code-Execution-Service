-- +migrate Up
CREATE TABLE tasks (
    task_id VARCHAR(100) PRIMARY KEY,
    status VARCHAR(100) NOT NULL,
    result TEXT
);

-- +migrate Down
DROP TABLE tasks;