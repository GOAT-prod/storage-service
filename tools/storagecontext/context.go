package storagecontext

import (
	"context"
	"storage-service/settings"

	"github.com/GOAT-prod/goatlogger"
)

type StorageContext struct {
	ctx    context.Context
	logger *goatlogger.Logger
}

func New() StorageContext {
	logger := goatlogger.New(settings.GetAppName())

	return StorageContext{
		ctx:    context.Background(),
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
