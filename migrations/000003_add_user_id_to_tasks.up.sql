ALTER TABLE tasks ADD COLUMN user_id INTEGER;

ALTER TABLE tasks
    ADD CONSTRAINT fk_tasks_user_id
        FOREIGN KEY (user_id) REFERENCES users(id)
            ON DELETE CASCADE;

CREATE INDEX idx_tasks_user_id ON tasks(user_id);
CREATE INDEX idx_tasks_user_id_deleted_at ON tasks(user_id, deleted_at);

