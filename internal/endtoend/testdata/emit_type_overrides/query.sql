-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY id;

-- name: CreateBook :one
INSERT INTO books (payload)
VALUES (sqlc.arg(payload))
RETURNING *;
