package handlers

import "net/http"

// PingHandler
//
//	@tags		ping
//	@summary	Проверка сервиса
//	@Success	200		"pong"
//	@Router		/ping [get]
func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}
}
