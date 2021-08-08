package helper

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/fauzanmh/online-store/constant"
	"github.com/lib/pq"
)

var pqErrorMap = map[string]int{
	"unique_violation": http.StatusConflict,
}

// PqError is
func PqError(err error) (int, error) {
	re := regexp.MustCompile("\\((.*?)\\)")
	if err, ok := err.(*pq.Error); ok {
		match := re.FindStringSubmatch(err.Detail)

		switch err.Code.Name() {
		case "unique_violation":
			return pqErrorMap["unique_violation"], fmt.Errorf("%s already exists", match[1])
		}
	}
	return http.StatusInternalServerError, fmt.Errorf("internal error")
}

var commonErrorMap = map[error]int{
	constant.ErrorPgUserAlreadyExists:      http.StatusConflict,
	constant.ErrorPgUserNotFound:           http.StatusNotFound,
	constant.ErrorPgDataNotFound:           http.StatusBadRequest,
	constant.ErrorMessageProductOutOfStock: http.StatusBadRequest,
}

// CommonError is
func CommonError(err error) (int, error) {
	switch err {
	case constant.ErrorPgUserAlreadyExists:
		return commonErrorMap[constant.ErrorPgUserAlreadyExists], constant.ErrorPgUserAlreadyExists
	case constant.ErrorPgUserNotFound:
		return commonErrorMap[constant.ErrorPgUserNotFound], constant.ErrorPgUserNotFound
	case constant.ErrorPgDataNotFound:
		return commonErrorMap[constant.ErrorPgDataNotFound], constant.ErrorPgDataNotFound
	case constant.ErrorMessageProductOutOfStock:
		return commonErrorMap[constant.ErrorMessageProductOutOfStock], constant.ErrorMessageProductOutOfStock
	}

	return http.StatusInternalServerError, fmt.Errorf("internal error")
}
