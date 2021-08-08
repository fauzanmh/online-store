package pg

import (
	"context"

	"github.com/fauzanmh/online-store/entity"
)

const storeOrder = `-- name: StoreOrder :one
INSERT INTO orders (user_id, created_at, updated_at)
VALUES ($1, $2, $3)
RETURNING id
`

func (q *Queries) StoreOrder(ctx context.Context, arg *entity.StoreOrderParams) (id int64, err error) {
	row := q.queryRow(ctx, q.storeOrderStmt, storeOrder, arg.UserID, arg.CreatedAt, arg.UpdatedAt)
	err = row.Scan(&id)
	return
}
