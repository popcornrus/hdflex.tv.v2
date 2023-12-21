package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MysqlInterface interface {
	DoInTransaction(do func(tx *sql.Tx) error) error
	Begin() error
	Commit() error
	Rollback() error
	GetTx() *sql.Tx
	GetExecer() QueryExecer
}

type MongoDBInterface interface {
	FindMongo(ctx context.Context, collection string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOneMongo(ctx context.Context, collection string, filter interface{}) *mongo.SingleResult
	InsertOneMongo(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error)
	UpdateOneMongo(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	DeleteOneMongo(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)
	DeleteManyMongo(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)
}

type QueryExecer interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type DataBase struct {
	db    *sql.DB
	tx    *sql.Tx
	mongo MongoConnection
}

func NewMysql(
	db *sql.DB,
) *DataBase {
	return &DataBase{
		db: db,
	}
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

func (u *DataBase) Begin() error {
	var err error
	u.tx, err = u.db.Begin()
	return err
}

func (u *DataBase) DoInTransaction(do func(tx *sql.Tx) error) (err error) {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = rollbackErr
				return
			}
			panic(p)
		} else if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = rollbackErr
				return
			}
		} else {
			err = tx.Commit()
		}
	}()

	err = do(tx)
	return err
}

func (u *DataBase) Commit() error {
	if u.tx == nil {
		return nil
	}

	err := u.tx.Commit()
	if err != nil {
		return err
	}

	u.tx = nil
	return err
}

func (u *DataBase) Rollback() error {
	if u.tx == nil {
		return fmt.Errorf("transaction is not started: %w", errors.New("rollback"))
	}

	err := u.tx.Rollback()
	if err != nil {
		return err
	}

	u.tx = nil
	return nil
}

func (u *DataBase) GetDB() *sql.DB {
	return u.db
}

func (u *DataBase) GetTx() *sql.Tx {
	return u.tx
}

func (u *DataBase) GetExecer() QueryExecer {
	if u.tx != nil {
		return u.tx
	}
	return u.db
}

func (u *DataBase) CheckConnection(db *sql.DB) error {
	conn, err := db.Conn(context.Background())
	if err != nil {
		return err
	}
	defer func(conn *sql.Conn) {
		if err := conn.Close(); err != nil {
			return
		}
	}(conn)

	if err = conn.PingContext(context.Background()); err != nil {
		return err
	}

	return nil
}
