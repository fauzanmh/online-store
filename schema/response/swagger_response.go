package response

type SwaggerProductsResponse struct {
	Base
	Data []ProductResponse `json:"data"`
}
