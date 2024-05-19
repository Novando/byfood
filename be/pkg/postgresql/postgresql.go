package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/novando/byfood/be/pkg/reposqlc"
)

func Init(
	user string,
	pass string,
	host string,
	port int,
	name string,
) (
	pool *pgxpool.Pool,
	query *reposqlc.Queries,
	err error,
) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?pool_max_conns=10",
		user,
		pass,
		host,
		port,
		name,
	)
	c, err := pgxpool.ParseConfig(url)
	if err != nil {
		return
	}

	pool, err = pgxpool.NewWithConfig(context.Background(), c)
	if err != nil {
		return
	}

	query = reposqlc.New(pool)
	return
}
