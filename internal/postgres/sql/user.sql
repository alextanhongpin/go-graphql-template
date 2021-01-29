-- name: CreateUser :one
INSERT INTO "user" (name, email, preferred_username)
VALUES ($1, $2, $3)
RETURNING *;

-- name: FindUser :one
SELECT *
FROM "user"
WHERE id = $1;

-- name: FindUsersWithIDs :many
SELECT *
FROM "user"
WHERE id = any($1::uuid[]);

-- name: UpdateUser :one
UPDATE "user"
SET name = COALESCE($1, name)
WHERE id = $2
RETURNING *;

-- name: DeleteUser :one
DELETE FROM "user"
WHERE id = $1
RETURNING *;
