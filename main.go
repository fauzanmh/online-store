package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/fauzanmh/online-store/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_orderHttp "github.com/fauzanmh/online-store/module/order/handler/http"
	_order "github.com/fauzanmh/online-store/module/order/usecase"
	_productHttp "github.com/fauzanmh/online-store/module/product/handler/http"
	_product "github.com/fauzanmh/online-store/module/product/usecase"
	_pgRepo "github.com/fauzanmh/online-store/repository/pg"

	appInit "github.com/fauzanmh/online-store/init"
	appMiddleware "github.com/fauzanmh/online-store/middleware"
	_ "github.com/spf13/viper/remote"
	echoSwagger "github.com/swaggo/echo-swagger"
	log "go.uber.org/zap"
)

var cfg *appInit.Config

func init() {
	// Start pre-requisite app dependencies
	cfg = appInit.StartAppInit()
}

func main() {
	// echo
	e := echo.New()

	// mutex
	var mtx sync.Mutex

	// timeout
	timeoutContext := time.Duration(cfg.Context.Timeout) * time.Second

	// init database
	pgDb, err := appInit.ConnectToPGServer(cfg)
	if err != nil {
		log.S().Fatal(err)
	}

	// init repository
	pgRepo := _pgRepo.NewRepository(pgDb)

	// init usecase
	// product usecase
	productUc := _product.NewProductUseCase(cfg, pgRepo)
	// order usecase
	orderUc := _order.NewOrderUseCase(cfg, &mtx, pgRepo)

	// Middleware
	e.Use(appMiddleware.EchoCORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(appMiddleware.DumpRequestResponse))
	config.SetEchoErrorDefault(e)
	// End of middleware

	// Grouping Routes
	routerAPI := e.Group("/api")
	// swagger route
	routerAPI.GET("/swagger/*", echoSwagger.WrapHandler)
	// order routes
	_orderHttp.NewOrderHandler(routerAPI, orderUc)
	// product routes
	_productHttp.NewProductHandler(routerAPI, productUc)

	go runHTTPHandler(e, cfg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func runHTTPHandler(e *echo.Echo, cfg *appInit.Config) {
	if err := e.Start(cfg.API.HTTP.Port); err != nil {
		fmt.Println("shutting down the server")
	}
}
