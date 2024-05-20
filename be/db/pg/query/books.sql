-- name: BookCount :one
SELECT COUNT(*) FROM books
WHERE
    title ILIKE '%'||sqlc.arg(title)::text||'%' AND
    yop = @yop::smallint;

-- name: BookCreate :exec
INSERT INTO books(title, yop, author, isbn, page)
VALUES (@title::text, @yop::smallint, @author::text, @isbn, @page);

-- name: BookUpdateById :exec
UPDATE books SET
    title = @title::text,
    yop = @yop::smallint,
    author = @author::text,
    isbn = @isbn,
    page = @page,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id::uuid;

-- name: BookDetailById :one
SELECT * FROM books WHERE id = @id::uuid;

-- name: BookDeleteById :exec
UPDATE books SET deleted_at = CURRENT_TIMESTAMP WHERE id = @id::uuid;