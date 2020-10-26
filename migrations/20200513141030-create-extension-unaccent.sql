-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "unaccent";

-- This is required for the slugify function.

-- +migrate Down
DROP EXTENSION IF EXISTS "unaccent";
