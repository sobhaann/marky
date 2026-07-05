-- name: SelectUsers :many
SELECT * FROM users
WHERE deleted_at IS NULL;

-- name: SelectUserByRID :one
SELECT * FROM users
WHERE r_id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: SelectUserByEmail :one
SELECT * FROM users
WHERE email = $1 AND deleted_at IS NULL
LIMIT 1;
