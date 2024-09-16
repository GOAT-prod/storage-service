package handlers

import (
	"fmt"
	"net/http"
	"storage-service/service"
	"storage-service/tools/goathttp"
	"storage-service/tools/storagecontext"
	"strconv"
)

func GetProducts(storageService service.StorageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storageCtx := storagecontext.New(r)
		storageCtx.SetLogTag("[get-products]")

		limit, err := parseLimit(r)
		if err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось распарсить лимит, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		products, err := storageService.GetProducts(storageCtx, limit)
		if err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось получить список продуктов, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = goathttp.WriteResponseJson(w, http.StatusOK, products); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось сериализовать список продуктов, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func parseLimit(r *http.Request) (int, error) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		return 0, nil
	}

	return strconv.Atoi(limit)
}
