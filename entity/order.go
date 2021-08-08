package entity

import (
	"database/sql"

	"github.com/google/uuid"
)

// model
type Order struct {
	ID        int64         `json:"id"`
	UserID    uuid.UUID     `json:"user_id"`
	Status    string        `json:"status"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt sql.NullInt64 `json:"updated_at"`
	DeletedAt sql.NullInt64 `json:"deleted_at"`
}

// params and rows
type StoreOrderParams struct {
	UserID    uuid.UUID     `json:"user_id"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt sql.NullInt64 `json:"updated_at"`
}
