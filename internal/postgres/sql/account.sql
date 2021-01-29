-- name: CreateAccount :one
INSERT INTO "account" (uid, provider, token, user_id, email)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: FindAccount :one
SELECT *
FROM "account"
WHERE id = $1;

-- name: FindAccountsWithIDs :many
SELECT *
FROM "account"
WHERE id = any($1::uuid[]);

-- name: FindAccountsWithUserID :many
SELECT *
FROM account
WHERE user_id = $1;

-- name: UpdateAccount :one
UPDATE "account"
SET email = COALESCE($1, email)
WHERE id = $2
RETURNING *;

-- name: DeleteAccount :one
DELETE FROM "account"
WHERE id = $1
RETURNING *;
