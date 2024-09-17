package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseInterface interface {
	Find(ctx context.Context, collection string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, collection string, filter interface{}) *mongo.SingleResult
	InsertOne(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)
}

type DataBase struct {
	mongo Connection
}

func NewMongo(
	mongoConnection Connection,
) *DataBase {
	return &DataBase{
		mongo: mongoConnection,
	}
}

func (u *DataBase) Find(ctx context.Context, collection string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	col := u.mongo.Client.Database(u.mongo.DbName).Collection(collection)
	return col.Find(ctx, filter, opts...)
}

func (u *DataBase) FindOne(ctx context.Context, collection string, filter interface{}) *mongo.SingleResult {
	col := u.mongo.Client.Database(u.mongo.DbName).Collection(collection)
	return col.FindOne(ctx, filter)
}

func (u *DataBase) InsertOne(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	col := u.mongo.Client.Database(u.mongo.DbName).Collection(collection)
	return col.InsertOne(ctx, document)
}

func (u *DataBase) UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	col := u.mongo.Client.Database(u.mongo.DbName).Collection(collection)
	return col.UpdateOne(ctx, filter, update)
}

func (u *DataBase) DeleteOne(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	col := u.mongo.Client.Database(u.mongo.DbName).Collection(collection)
	return col.DeleteOne(ctx, filter)
}

func (u *DataBase) DeleteMany(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	col := u.mongo.Client.Database(u.mongo.DbName).Collection(collection)
	return col.DeleteMany(ctx, filter)
}
