package http

import (
	"github.com/fauzanmh/online-store/module/product/usecase"
	"github.com/fauzanmh/online-store/pkg/util"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	usecase usecase.Usecase
}

func NewProductHandler(e *echo.Group, uc usecase.Usecase) {
	handler := &ProductHandler{
		usecase: uc,
	}

	// router for product group
	routerV1 := e.Group("/v1")
	routerV1.GET("/products", handler.GetAllProduct)
}

// GetAllProduct godoc
// @Summary Get All Product
// @Description Endpoint for get all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} response.SwaggerProductsResponse
// @Failure 400 {object} response.Base
// @Failure 401 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/products [get]
func (h *ProductHandler) GetAllProduct(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := h.usecase.GetAllProducts(ctx)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success get products", data)
}
