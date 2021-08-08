package pg

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.getAllProductsStmt, err = db.PrepareContext(ctx, getAllProducts); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllProducts: %w", err)
	}
	if q.getProductStmt, err = db.PrepareContext(ctx, getProduct); err != nil {
		return nil, fmt.Errorf("error preparing query GetProduct: %w", err)
	}
	if q.storeOrderStmt, err = db.PrepareContext(ctx, storeOrder); err != nil {
		return nil, fmt.Errorf("error preparing query StoreOrder: %w", err)
	}
	if q.storeOrderDetailStmt, err = db.PrepareContext(ctx, storeOrderDetail); err != nil {
		return nil, fmt.Errorf("error preparing query StoreOrderDetail: %w", err)
	}
	if q.updateStockProductStmt, err = db.PrepareContext(ctx, updateStockProduct); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateStockProduct: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.getAllProductsStmt != nil {
		if cerr := q.getAllProductsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllProductsStmt: %w", cerr)
		}
	}
	if q.getProductStmt != nil {
		if cerr := q.getProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductStmt: %w", cerr)
		}
	}
	if q.storeOrderStmt != nil {
		if cerr := q.storeOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing storeOrderStmt: %w", cerr)
		}
	}
	if q.storeOrderDetailStmt != nil {
		if cerr := q.storeOrderDetailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing storeOrderDetailStmt: %w", cerr)
		}
	}
	if q.updateStockProductStmt != nil {
		if cerr := q.updateStockProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateStockProductStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                     DBTX
	tx                     *sql.Tx
	getAllProductsStmt     *sql.Stmt
	getProductStmt         *sql.Stmt
	storeOrderStmt         *sql.Stmt
	storeOrderDetailStmt   *sql.Stmt
	updateStockProductStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                     tx,
		tx:                     tx,
		getAllProductsStmt:     q.getAllProductsStmt,
		getProductStmt:         q.getProductStmt,
		storeOrderStmt:         q.storeOrderStmt,
		storeOrderDetailStmt:   q.storeOrderDetailStmt,
		updateStockProductStmt: q.updateStockProductStmt,
	}
}
