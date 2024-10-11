package database

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"storage-service/tools/storagecontext"
	"time"
)

type LogItem struct {
	Id        primitive.ObjectID `bson:"_id"`
	Method    string             `bson:"method"`
	EventDate string             `bson:"time"`
	Data      string             `bson:"data"`
}

type LogsRepository interface {
	Log(ctx storagecontext.StorageContext, method string, data any) error
}

type LogsRepositoryImpl struct {
	mongoClient *mongo.Client
	database    string
	collection  string
}

func NewLogsRepository(mongoClient *mongo.Client, database, collection string) LogsRepository {
	return &LogsRepositoryImpl{
		mongoClient: mongoClient,
		database:    database,
		collection:  collection,
	}
}

func (r *LogsRepositoryImpl) Log(ctx storagecontext.StorageContext, method string, data any) error {
	dataB, err := json.Marshal(data)
	if err != nil {
		return err
	}

	logItem := LogItem{
		Id:        primitive.NewObjectID(),
		Method:    method,
		EventDate: time.Now().Format(time.RFC3339),
		Data:      string(dataB),
	}

	if _, err = r.mongoClient.Database(r.database).Collection(r.collection).InsertOne(ctx.Ctx(), logItem); err != nil {
		return err
	}

	return nil
}
