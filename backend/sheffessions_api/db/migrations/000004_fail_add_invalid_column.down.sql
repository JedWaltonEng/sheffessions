-- Since the migration will fail, you may not need to revert this in the .down.sql
-- But just for completeness:
ALTER TABLE articles DROP COLUMN IF EXISTS invalid_column;

