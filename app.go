package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/GOAT-prod/goatlogger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"storage-service/api"
	"storage-service/database"
	"storage-service/service"
	"storage-service/settings"
	"time"
)

type App struct {
	mainCtx context.Context
	config  settings.Config
	logger  goatlogger.Logger

	server *http.Server

	storageService service.StorageService

	mongo             *mongo.Client
	postgres          *sqlx.DB
	storageRepository database.StorageRepository
	logsRepository    database.LogsRepository
}

func NewApp(ctx context.Context, config settings.Config, logger goatlogger.Logger) *App {
	return &App{
		mainCtx: ctx,
		config:  config,
		logger:  logger,
	}
}

func (a *App) Start() {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error(fmt.Sprintf("приложение неожиданно остановлено, ошибка: %v", err))
		}
	}()
}

func (a *App) Stop(ctx context.Context) {
	stopCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	if err := a.server.Shutdown(stopCtx); err != nil {
		a.logger.Error(fmt.Sprintf("Не удалось остановить сервер: %v", err))
	}

	if err := a.mongo.Disconnect(stopCtx); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось отключиться от монги: %v", err))
	}

	if err := a.postgres.Close(); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось закрыть подключение к постгресу: %v", err))
	}
}

func (a *App) initDatabases() {
	a.initPostgres()
	a.initMongo()
}

func (a *App) initPostgres() {
	a.postgres = sqlx.MustConnect("postgres", a.config.Databases.Postgres)

	if err := database.RunMigrations(a.postgres, "./database/migrations"); err != nil {
		a.logger.Panic(fmt.Sprintf("не удалось применить миграции, ошибка: %v", err))
		os.Exit(1)
	}
}

func (a *App) initMongo() {
	mongoCtx, cancelFunc := context.WithTimeout(a.mainCtx, 15*time.Second)
	defer cancelFunc()

	mongoClient, err := database.MongoConnect(mongoCtx, a.config.Databases.MongoDB.ConnectionString)
	if err != nil {
		a.logger.Panic(fmt.Sprintf("не удалось подключиться к mongoDb, ошибка: %v", err))
		os.Exit(1)
	}

	a.mongo = mongoClient
}

func (a *App) initRepositories() {
	a.storageRepository = database.NewStorageRepository(a.postgres)

	if settings.GetEnv() == settings.LocalEnv() && a.config.Databases.NeedMocks {
		if err := a.storageRepository.InsertForTest(); err != nil {
			a.logger.Error(err.Error())
		}
	}

	a.logsRepository = database.NewLogsRepository(a.mongo, a.config.Databases.MongoDB.Database, a.config.Databases.MongoDB.Collection)
}

func (a *App) initServices() {
	a.storageService = service.NewStorageService(a.storageRepository, a.logsRepository)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Error("сервер уже запущен")
		os.Exit(1)
	}

	a.server = api.NewServer(a.mainCtx, a.logger, a.config, a.storageService)
}
