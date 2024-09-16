package storagecontext

import (
	"context"
	"net/http"
	"storage-service/settings"

	"github.com/GOAT-prod/goatlogger"
)

type StorageContext struct {
	ctx    context.Context
	logger *goatlogger.Logger
}

func New(r *http.Request) StorageContext {
	logger := goatlogger.New(settings.GetAppName())

	return StorageContext{
		ctx:    r.Context(),
		logger: &logger,
	}
}

func (sc *StorageContext) SetCtx(ctx context.Context) {
	sc.ctx = ctx
}

func (sc *StorageContext) Ctx() context.Context {
	return sc.ctx
}

func (sc *StorageContext) Log() *goatlogger.Logger {
	return sc.logger
}

func (sc *StorageContext) SetLogTag(tag string) {
	sc.logger.SetTag(tag)
}
