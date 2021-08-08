package usecase

import (
	"context"

	"github.com/fauzanmh/online-store/schema/response"
)

type Usecase interface {
	GetAllProducts(ctx context.Context) (res []response.ProductResponse, err error)
}
