package usecase

import (
	"context"

	"github.com/fauzanmh/online-store/schema/request"
)

type Usecase interface {
	Checkout(ctx context.Context, req *request.CheckoutRequest) (err error)
}
