-- name: BookCreate :exec
INSERT INTO books(title, yop, author, isbn, page)
VALUES (@title::text, @yop::smallint, @author::text, @isbn::text, @page::int);

-- name: BookUpdateById :exec
UPDATE books SET
    title = @title::text,
    yop = @yop::smallint,
    author = @author::text,
    isbn = @isbn::text,
    page = @page::int,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id::uuid;

-- name: BookDeleteById :exec
UPDATE books SET deleted_at = CURRENT_TIMESTAMP WHERE id = @id::uuid;