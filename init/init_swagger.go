package init

import "github.com/fauzanmh/online-store/docs"

func setupSwagger() {
	docs.SwaggerInfo.Title = "Online Store API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8099"
	docs.SwaggerInfo.BasePath = "/api"
}
