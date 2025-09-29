DROP INDEX IF EXISTS idx_tasks_user_id;
DROP INDEX IF EXISTS idx_tasks_user_id_deleted_at;

ALTER TABLE tasks DROP CONSTRAINT fk_tasks_user_id;

ALTER TABLE tasks DROP COLUMN user_id;