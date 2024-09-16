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
	router.Use(goathttp.CommonJsonMiddleware, goathttp.CommonJsonMiddleware, goathttp.PanicRecoveryMiddleware(logger))

	addProductHandlers(router, storageService)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
		Handler: router,
	}
}

func addProductHandlers(router *mux.Router, storageService service.StorageService) {
	router.HandleFunc("/products", handlers.GetProducts(storageService)).Methods(http.MethodGet)
}
