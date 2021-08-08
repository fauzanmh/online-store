package response

type ProductResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
	Stock int32  `json:"stock"`
}
