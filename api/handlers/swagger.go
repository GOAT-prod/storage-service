package handlers

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
	_ "storage-service/docs"
	"storage-service/settings"
)

func SwaggerHandler(settings settings.Config) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", settings.Port)), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"))
}
