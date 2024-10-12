package handlers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"storage-service/service"
	"storage-service/tools/storagecontext"
	"strconv"
)

// DeleteProductHandler
//
//	@tags		products
//	@summary	Удаление продукта
//	@accept		json
//	@produce	json
//	@security	Bearer
//	@param		productId	path		int		true	"ID продукта для удаления"
//	@Success	200			"Продукт успешно удален"
//	@Failure	400			"Ошибка запроса"
//	@Failure	401			"Не авторизован"
//	@Router		/product/{productId} [delete]
func DeleteProductHandler(storageService service.StorageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storageCtx := storagecontext.New(r)
		storageCtx.SetLogTag("[delete-product]")

		if !storageCtx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		productId, err := parseProductId(r)
		if err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось распарсить id продукта, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = storageService.RemoveProduct(storageCtx, productId); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось удалить продукт, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func parseProductId(r *http.Request) (int, error) {
	routeProductId, ok := mux.Vars(r)["productId"]
	if !ok {
		return 0, errors.New("отсутствует id продукта в пути запроса")
	}

	productId, err := strconv.Atoi(routeProductId)
	if err != nil {
		return 0, fmt.Errorf("не валидный id продукта: %w", err)
	}

	return productId, nil
}
