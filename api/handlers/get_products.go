package handlers

import (
	"fmt"
	"net/http"
	"storage-service/service"
	"storage-service/tools/goathttp"
	"storage-service/tools/storagecontext"
	"strconv"
)

// GetProductsHandler
//
//	@tags		products
//	@summary	Получение списка продуктов
//	@accept		json
//	@produce	json
//	@security	Bearer
//	@param		limit		query		int		true	"Количество продуктов на странице"
//	@param		page		query		int		true	"Номер страницы для пагинации"
//	@Success	200			{array}		domain.Product	"Успешное получение списка продуктов"
//	@Failure	400			"Ошибка запроса"
//	@Router		/products [get]
func GetProductsHandler(storageService service.StorageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storageCtx := storagecontext.New(r)
		storageCtx.SetLogTag("[get-products]")

		limit, page := parseQuery(r)

		products, err := storageService.GetProducts(storageCtx, limit, page)
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

func parseQuery(r *http.Request) (int, int) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	return limit, page
}
