package api

import (
	"context"
	"fmt"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"storage-service/api/handlers"
	"storage-service/service"
	"storage-service/settings"
	"storage-service/tools/goathttp"
)

func NewServer(ctx context.Context, logger goatlogger.Logger, cfg settings.Config, storageService service.StorageService) *http.Server {
	router := mux.NewRouter()
	router.Use(goathttp.CommonJsonMiddleware, goathttp.CORSMiddleware, goathttp.PanicRecoveryMiddleware(logger))

	router.HandleFunc("/ping", handlers.PingHandler()).Methods(http.MethodGet)

	addProductHandlers(router, storageService)
	addSwaggerHandler(router, cfg)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
		Handler: router,
	}
}

func addProductHandlers(router *mux.Router, storageService service.StorageService) {
	router.HandleFunc("/products", handlers.GetProductsHandler(storageService)).Methods(http.MethodGet)
	router.HandleFunc("/products", handlers.AddProductHandler(storageService)).Methods(http.MethodPost)
	router.HandleFunc("/products", handlers.UpdateProductHandler(storageService)).Methods(http.MethodPut)
	router.HandleFunc("/product/{productId}", handlers.DeleteProductHandler(storageService)).Methods(http.MethodDelete)
}

func addSwaggerHandler(router *mux.Router, cfg settings.Config) {
	router.PathPrefix("/swagger/").Handler(handlers.SwaggerHandler(cfg))
}
