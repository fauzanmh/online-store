package util

import (
	"os"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

//ToSnakeCase is function to convert camelCase to snake_case
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

//SetLowerAndAddSpace is function to convert camelCase to snake_case
func SetLowerAndAddSpace(str string) string {
	lower := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	lower = matchAllCap.ReplaceAllString(lower, "${1} ${2}")
	return strings.ToLower(lower)
}

// GetEnv returns app envorinment : e.g. development, production, staging, testing, etc
func GetEnv() string {
	return os.Getenv("APP_ENV")
}

// IsProductionEnv returns whether the app is running using production env
func IsProductionEnv() bool {
	return os.Getenv("APP_ENV") == "production"
}

// IsDevelopmentEnv returns whether the app is running using production env
func IsDevelopmentEnv() bool {
	return os.Getenv("APP_ENV") == "development"
}
