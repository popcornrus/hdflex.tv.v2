package db

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBInterface interface {
	FindMongo(ctx context.Context, collection string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOneMongo(ctx context.Context, collection string, filter interface{}) *mongo.SingleResult
	InsertOneMongo(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error)
	UpdateOneMongo(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	DeleteOneMongo(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)
	DeleteManyMongo(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)
}

type DataBase struct {
	mongo MongoConnection
}

func NewMongo(
	mongoConnection MongoConnection,
) *DataBase {
	return &DataBase{
		mongo: mongoConnection,
	}
}

func (u *DataBase) FindMongo(ctx context.Context, collection string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	col := u.mongo.MongoClient.Database(u.mongo.DBName).Collection(collection)
	return col.Find(ctx, filter, opts...)
}

func (u *DataBase) FindOneMongo(ctx context.Context, collection string, filter interface{}) *mongo.SingleResult {
	col := u.mongo.MongoClient.Database(u.mongo.DBName).Collection(collection)
	return col.FindOne(ctx, filter)
}

func (u *DataBase) InsertOneMongo(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	col := u.mongo.MongoClient.Database(u.mongo.DBName).Collection(collection)
	return col.InsertOne(ctx, document)
}

func (u *DataBase) UpdateOneMongo(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	col := u.mongo.MongoClient.Database(u.mongo.DBName).Collection(collection)
	return col.UpdateOne(ctx, filter, update)
}

func (u *DataBase) DeleteOneMongo(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	col := u.mongo.MongoClient.Database(u.mongo.DBName).Collection(collection)
	return col.DeleteOne(ctx, filter)
}

func (u *DataBase) DeleteManyMongo(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	col := u.mongo.MongoClient.Database(u.mongo.DBName).Collection(collection)
	return col.DeleteMany(ctx, filter)
}
