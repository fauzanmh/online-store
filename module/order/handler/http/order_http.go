package http

import (
	"github.com/fauzanmh/online-store/module/order/usecase"
	"github.com/fauzanmh/online-store/pkg/util"
	"github.com/fauzanmh/online-store/schema/request"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	usecase usecase.Usecase
}

func NewOrderHandler(e *echo.Group, uc usecase.Usecase) {
	handler := &OrderHandler{
		usecase: uc,
	}

	// router for order group
	routerV1 := e.Group("/v1")
	routerV1.POST("/orders/checkout", handler.OrderCheckout)
}

// OrderCheckout godoc
// @Summary Checkout
// @Description Checkout for order *user_id auto generate
// @Tags Order
// @Accept json
// @Produce json
// @Param request body request.CheckoutRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 401 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/orders/checkout [post]
func (h *OrderHandler) OrderCheckout(c echo.Context) error {
	req := request.CheckoutRequest{}
	ctx := c.Request().Context()

	// parsing
	err := util.ParsingParameter(c, &req)
	if err != nil {
		return util.ErrorParsing(c, err, nil)
	}

	// validate
	err = util.ValidateParameter(c, &req)
	if err != nil {
		return util.ErrorValidate(c, err, nil)
	}

	err = h.usecase.Checkout(ctx, &req)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success create order", nil)
}
