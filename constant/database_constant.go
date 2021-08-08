package constant

import "fmt"

type ErrorPg error

var (
	ErrorPgUserNotFound      ErrorPg = fmt.Errorf("user not found")
	ErrorPgUserAlreadyExists ErrorPg = fmt.Errorf("user already exists")
	ErrorPgDataNotFound      ErrorPg = fmt.Errorf("data not found")
)
