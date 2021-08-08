package entity

import "database/sql"

// model
type Product struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Price     string        `json:"price"`
	Stock     int32         `json:"stock"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt sql.NullInt64 `json:"updated_at"`
}

// params and row
type GetAllProductsRow struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Stock int32  `json:"stock"`
	Price string `json:"price"`
}

type GetProductRow struct {
	Stock int32  `json:"stock"`
	Price string `json:"price"`
}

type UpdateStockProductParams struct {
	Stock     int32         `json:"stock"`
	UpdatedAt sql.NullInt64 `json:"updated_at"`
	ID        int64         `json:"id"`
}
