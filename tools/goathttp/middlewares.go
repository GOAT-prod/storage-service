package goathttp

import (
	"fmt"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

func PanicRecoveryMiddleware(logger goatlogger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		logger.SetTag("[PANIC RECOVERY]")
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				panicErr := recover()
				if panicErr != nil {
					logger.Panic(fmt.Sprintf("произошла паника: %s", panicErr))
					jsonPanic := map[string]any{
						"panic": "что-то пошло не так",
					}
					_ = WriteResponseJson(w, http.StatusInternalServerError, jsonPanic)
				}

			}()
			next.ServeHTTP(w, r)
		})
	}
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(_accessControlAllowOriginHeader, _allowedOrigins)
		w.Header().Add(_accessControlAllowMethodsHeader, _allowedMethods)
		w.Header().Add(_accessControlAllowHeaders, _allowedHeaders)
		w.Header().Add(_accessControlAllowsCredentialsHeader, "true")

		next.ServeHTTP(w, r)
	})
}

func OptionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (*r).Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CommonJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
