package handlers

import (
	"fmt"
	"net/http"
	"storage-service/domain"
	"storage-service/service"
	"storage-service/tools/goathttp"
	"storage-service/tools/storagecontext"
)

// AddProductHandler
//
//	@tags		products
//	@summary	Добавление нового продукта
//	@accept		json
//	@produce	json
//	@security	Bearer
//	@param		Product	body		domain.Product	true	"Данные нового продукта"
//	@Success	200		"Продукт успешно добавлен"
//	@Failure	400		"Ошибка запроса"
//	@Router		/products [post]
func AddProductHandler(storageService service.StorageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storageCtx := storagecontext.New(r)
		storageCtx.SetLogTag("[add-product]")

		var product domain.Product
		if err := goathttp.ReadRequestJson(r, &product); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось распарсить новый продукт, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := storageService.SaveProduct(storageCtx, product); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось добавить новый продукт, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
