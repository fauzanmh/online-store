package pg

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/online-store/entity"
)

type Repository interface {
	// Product
	GetAllProducts(ctx context.Context) (res []entity.GetAllProductsRow, err error)
	GetProduct(ctx context.Context, id int64) (res entity.GetProductRow, err error)
	UpdateStockProduct(ctx context.Context, arg *entity.UpdateStockProductParams) (err error)

	// Order
	StoreOrder(ctx context.Context, arg *entity.StoreOrderParams) (id int64, err error)

	// Order Detail
	StoreOrderDetail(ctx context.Context, arg *entity.StoreOrderDetailParams) (err error)

	//Tx
	BeginTx(ctx context.Context) (*sql.Tx, error)
	WithTx(tx *sql.Tx) *Queries
	RollbackTx(tx *sql.Tx) error
	CommitTx(tx *sql.Tx) error
}
