package handlers

import (
	"fmt"
	"net/http"
	"storage-service/domain"
	"storage-service/service"
	"storage-service/tools/goathttp"
	"storage-service/tools/storagecontext"
)

// UpdateProductHandler
//
//	@tags		products
//	@summary	Обновление информации о продукте
//	@accept		json
//	@produce	json
//	@security	Bearer
//	@param		Product	body		domain.Product	true	"Данные обновляемого продукта"
//	@Success	200		"Продукт успешно обновлен"
//	@Failure	400		"Ошибка запроса"
//	@Failure	401		"Не авторизован"
//	@Router		/products [put]
func UpdateProductHandler(storageService service.StorageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storageCtx := storagecontext.New(r)
		storageCtx.SetLogTag("[update-product]")

		if !storageCtx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var product domain.Product
		if err := goathttp.ReadRequestJson(r, &product); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось распарсить обновленный продукт, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := storageService.UpdateProduct(storageCtx, product); err != nil {
			storageCtx.Log().Error(fmt.Sprintf("не удалось обновить продукт, ошибка: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
