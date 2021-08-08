package constant

import "fmt"

type ErrorMessage error

var (
	ErrorMessageProductOutOfStock ErrorMessage = fmt.Errorf("product out of stock")
)
