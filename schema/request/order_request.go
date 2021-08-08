package request

type CheckoutRequest struct {
	Items []ItemsRequest `json:"items" validate:"gt=0,dive"`
}

type ItemsRequest struct {
	ProductID int64 `json:"product_id" validate:"required"`
	Qty       int32 `json:"qty" validate:"required"`
}
