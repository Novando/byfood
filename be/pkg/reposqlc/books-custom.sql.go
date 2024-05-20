package reposqlc

import (
	"context"
	"fmt"
)

const bookGetAll = `-- name: BookGetAll :many
SELECT id, title, yop, author, isbn, page, created_at, updated_at, deleted_at FROM books
WHERE
    title = $1::text AND
    yop = $2::smallint
`

type BookGetAllParams struct {
	Title string
	Yop   int16
}

func (q *Queries) BookGetAll(
	ctx context.Context,
	xtra ColumnCustomParams,
	arg BookGetAllParams,
) ([]Book, error) {
	sorter := DESC
	if xtra.Ascending {
		sorter = ASC
	}
	query := fmt.Sprintf(
		"%v ORDER BY %v %v LIMIT %v OFFSET %v",
		bookGetAll,
		xtra.ColumnName,
		sorter,
		xtra.Limit,
		xtra.Offset,
	)
	rows, err := q.db.Query(ctx, query, arg.Title, arg.Yop)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Yop,
			&i.Author,
			&i.Isbn,
			&i.Page,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}