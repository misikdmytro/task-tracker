CREATE TABLE IF NOT EXISTS tbl_lists (
    id SERIAL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS tbl_tasks (
    id SERIAL,
    name VARCHAR(255) NOT NULL,
    list_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (list_id) REFERENCES tbl_lists(id)
);
CREATE INDEX IF NOT EXISTS idx_tasks_list_id ON tbl_tasks (list_id);