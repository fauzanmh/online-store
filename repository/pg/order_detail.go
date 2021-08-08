package pg

import (
	"context"

	"github.com/fauzanmh/online-store/entity"
)

const storeOrderDetail = `-- name: StoreOrderDetail :exec
INSERT INTO order_details (id, order_id, product_id, price, qty, total_price, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

func (q *Queries) StoreOrderDetail(ctx context.Context, arg *entity.StoreOrderDetailParams) (err error) {
	_, err = q.exec(ctx, q.storeOrderDetailStmt, storeOrderDetail,
		arg.ID,
		arg.OrderID,
		arg.ProductID,
		arg.Price,
		arg.Qty,
		arg.TotalPrice,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return
}
