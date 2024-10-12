package handlers

import (
	"fmt"
	"net/http"
	"storage-service/service"
	"storage-service/tools/goathttp"
	"storage-service/tools/storagecontext"
)

// GetProductHandler
//
//	@tags		products
//	@summary	Получение информации о продукте
//	@description Метод получения информации о конкретном продукте по его ID. Требуется авторизация.
//	@accept		json
//	@produce	json
//	@security	Bearer
//	@param		productId	path		integer		true	"ID продукта"
//	@Success	200			{object}	domain.Product	"Продукт успешно получен"
//	@Failure	400			"Ошибка запроса, не удалось получить продукт"
//	@Failure	401			"Пользователь не авторизован"
//	@Router		/product/{productId} [get]
func GetProductHandler(storageService service.StorageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storageCtx := storagecontext.New(r)
		storageCtx.SetLogTag("[get-products]")

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

		product, err := storageService.GetProduct(storageCtx, productId)
		if err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось получить продукт %d, ошибка: %v", productId, err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = goathttp.WriteResponseJson(w, http.StatusOK, product); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось сериализовать продукт, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
