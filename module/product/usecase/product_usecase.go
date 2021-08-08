package usecase

import (
	"context"

	appInit "github.com/fauzanmh/online-store/init"
	pgRepo "github.com/fauzanmh/online-store/repository/pg"
	"github.com/fauzanmh/online-store/schema/response"
	"go.uber.org/zap"
)

type usecase struct {
	config *appInit.Config
	pgRepo pgRepo.Repository
}

func NewProductUseCase(config *appInit.Config, pgRepo pgRepo.Repository) Usecase {
	return &usecase{
		config: config,
		pgRepo: pgRepo,
	}
}

func (u *usecase) GetAllProducts(ctx context.Context) (res []response.ProductResponse, err error) {
	// get data products from repository
	data, err := u.pgRepo.GetAllProducts(ctx)
	if err != nil {
		zap.S().Named("get.products").Error(err)
		return
	}

	for _, v := range data {
		res = append(res, response.ProductResponse{
			ID:    v.ID,
			Name:  v.Name,
			Price: v.Price,
			Stock: v.Stock,
		})
	}

	return
}
