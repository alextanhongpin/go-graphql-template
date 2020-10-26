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
SET name = COALESCE($1, name),
given_name = COALESCE($2, given_name),
family_name = COALESCE($3, family_name),
profile = COALESCE($4, profile)
WHERE id = $5
RETURNING *;
