package pg

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/online-store/constant"
	"github.com/fauzanmh/online-store/entity"
)

const getAllProducts = `-- name: GetAllProducts :many
SELECT id, name, stock, price 
FROM products
`

func (q *Queries) GetAllProducts(ctx context.Context) (res []entity.GetAllProductsRow, err error) {
	rows, err := q.query(ctx, q.getAllProductsStmt, getAllProducts)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var i entity.GetAllProductsRow
		if err = rows.Scan(
			&i.ID,
			&i.Name,
			&i.Stock,
			&i.Price,
		); err != nil {
			return
		}
		res = append(res, i)
	}
	if err = rows.Close(); err != nil {
		return
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

const getProduct = `-- name: GetProduct :one
SELECT stock, price 
FROM products
WHERE id = $1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (res entity.GetProductRow, err error) {
	row := q.queryRow(ctx, q.getProductStmt, getProduct, id)
	err = row.Scan(&res.Stock, &res.Price)
	if err == sql.ErrNoRows {
		err = constant.ErrorPgDataNotFound
	}
	return
}

const updateStockProduct = `-- name: UpdateStockProduct :exec
UPDATE products 
SET stock = $1, updated_at = $2
WHERE id = $3
`

func (q *Queries) UpdateStockProduct(ctx context.Context, arg *entity.UpdateStockProductParams) (err error) {
	_, err = q.exec(ctx, q.updateStockProductStmt, updateStockProduct, arg.Stock, arg.UpdatedAt, arg.ID)
	return
}
