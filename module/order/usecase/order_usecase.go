package usecase

import (
	"context"
	"database/sql"
	"strconv"
	"sync"
	"time"

	"github.com/fauzanmh/online-store/constant"
	"github.com/fauzanmh/online-store/entity"
	appInit "github.com/fauzanmh/online-store/init"
	pgRepo "github.com/fauzanmh/online-store/repository/pg"
	"github.com/fauzanmh/online-store/schema/request"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type usecase struct {
	config *appInit.Config
	mtx    *sync.Mutex
	pgRepo pgRepo.Repository
}

func NewOrderUseCase(config *appInit.Config, mtx *sync.Mutex, pgRepo pgRepo.Repository) Usecase {
	return &usecase{
		config: config,
		mtx:    mtx,
		pgRepo: pgRepo,
	}
}

func (u *usecase) Checkout(ctx context.Context, req *request.CheckoutRequest) (err error) {
	// lock process
	u.mtx.Lock()

	// begin transactions
	tx, err := u.pgRepo.BeginTx(ctx)
	if err != nil {
		return
	}

	// store order
	timeNowUnix := time.Now().Unix()
	storeOrderParams := entity.StoreOrderParams{
		UserID:    uuid.New(),
		CreatedAt: timeNowUnix,
		UpdatedAt: sql.NullInt64{
			Int64: timeNowUnix,
			Valid: true,
		},
	}

	orderID, err := u.pgRepo.WithTx(tx).StoreOrder(ctx, &storeOrderParams)
	if err != nil {
		zap.S().Named("store.order").Error(err)
		return
	}

	hasError := 0
	for idx := range req.Items {

		// get data product by id from repository
		product := entity.GetProductRow{}
		product, err = u.pgRepo.WithTx(tx).GetProduct(ctx, req.Items[idx].ProductID)
		if err != nil {
			hasError += 1
			break
		}

		// check product quantity
		if product.Stock < 1 || req.Items[idx].Qty > product.Stock {
			err = constant.ErrorMessageProductOutOfStock
			hasError += 1
			break
		}

		// store order detail
		// parse string to int
		price := 0
		price, err = strconv.Atoi(product.Price)
		if err != nil {
			hasError += 1
			break
		}

		// calculate total price
		totalPrice := price * int(req.Items[idx].Qty)
		storeOrderDetailParams := entity.StoreOrderDetailParams{
			ID:         int32(idx + 1),
			OrderID:    orderID,
			ProductID:  req.Items[idx].ProductID,
			Price:      product.Price,
			Qty:        req.Items[idx].Qty,
			TotalPrice: strconv.Itoa(totalPrice),
			CreatedAt:  timeNowUnix,
			UpdatedAt: sql.NullInt64{
				Int64: timeNowUnix,
				Valid: true,
			},
		}

		err = u.pgRepo.WithTx(tx).StoreOrderDetail(ctx, &storeOrderDetailParams)
		if err != nil {
			hasError += 1
			break
		}

		// update stock product after success store order
		updateStockProductParams := entity.UpdateStockProductParams{
			ID:    req.Items[idx].ProductID,
			Stock: product.Stock - req.Items[idx].Qty,
			UpdatedAt: sql.NullInt64{
				Int64: timeNowUnix,
				Valid: true,
			},
		}

		err = u.pgRepo.WithTx(tx).UpdateStockProduct(ctx, &updateStockProductParams)
		if err != nil {
			hasError += 1
			break
		}

	}

	// check transactions has error
	if hasError > 0 {
		u.pgRepo.RollbackTx(tx)
	} else {
		u.pgRepo.CommitTx(tx)
	}

	// unlock process
	u.mtx.Unlock()

	return
}
