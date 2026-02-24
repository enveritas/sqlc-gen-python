-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY title;

-- name: CreateBook :one
INSERT INTO books (
          title, status, payload, metadata
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: SearchBooks :many
SELECT * FROM books
WHERE payload @> COALESCE(sqlc.narg('payload_filter')::jsonb, '{}'::jsonb)
ORDER BY title;
