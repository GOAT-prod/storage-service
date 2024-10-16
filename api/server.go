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
	router.Use(goathttp.CommonJsonMiddleware, goathttp.CORSMiddleware, goathttp.OptionsMiddleware, goathttp.PanicRecoveryMiddleware(logger))

	router.HandleFunc("/ping", handlers.PingHandler()).Methods(http.MethodGet)

	addProductHandlers(router, storageService)
	addSwaggerHandler(router)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
		Handler: router,
	}
}

func addProductHandlers(router *mux.Router, storageService service.StorageService) {
	router.HandleFunc("/product/{productId}", handlers.GetProductHandler(storageService)).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/products", handlers.GetProductsHandler(storageService)).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/products", handlers.AddProductHandler(storageService)).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/products", handlers.UpdateProductHandler(storageService)).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/product/{productId}", handlers.DeleteProductHandler(storageService)).Methods(http.MethodDelete, http.MethodOptions)
}

func addSwaggerHandler(router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(handlers.SwaggerHandler())
}
