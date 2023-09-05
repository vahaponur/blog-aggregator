-- +goose Up
-- First, add a NOT NULL constraint to the user_id column
ALTER TABLE feeds
    ALTER COLUMN user_id SET NOT NULL;

-- +goose Down
-- To undo the change, you can remove the NOT NULL constraint
-- Note: If there are existing rows with NULL values in user_id, you should update them before removing the constraint.
-- You might want to consider a default value for new rows before setting it to NOT NULL.
ALTER TABLE feeds
    ALTER COLUMN user_id DROP NOT NULL;