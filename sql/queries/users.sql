-- name : Create : one
INSERT INTO users (id, created_at, modeified_at, name )
VALUES ( $1, $2, $3, $4)
RETURNING *;