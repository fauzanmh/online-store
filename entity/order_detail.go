package entity

import "database/sql"

// model
type OrderDetail struct {
	ID         int32         `json:"id"`
	OrderID    int64         `json:"order_id"`
	ProductID  int64         `json:"product_id"`
	Price      string        `json:"price"`
	Qty        int32         `json:"qty"`
	TotalPrice string        `json:"total_price"`
	CreatedAt  int64         `json:"created_at"`
	UpdatedAt  sql.NullInt64 `json:"updated_at"`
	DeletedAt  sql.NullInt64 `json:"deleted_at"`
}

// params and rows
type StoreOrderDetailParams struct {
	ID         int32         `json:"id"`
	OrderID    int64         `json:"order_id"`
	ProductID  int64         `json:"product_id"`
	Price      string        `json:"price"`
	Qty        int32         `json:"qty"`
	TotalPrice string        `json:"total_price"`
	CreatedAt  int64         `json:"created_at"`
	UpdatedAt  sql.NullInt64 `json:"updated_at"`
}
