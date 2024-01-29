package db

import (
	"context"
	"fmt"
	"go-chat/external/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"log/slog"
	"net/url"
)

type MongoConnection struct {
	MongoClient *mongo.Client
	DBName      string
}

func NewMongoDatabase(lc fx.Lifecycle, log *slog.Logger, cfg *config.Config) (MongoConnection, error) {
	var mongoConnection MongoConnection

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authMechanism=%s&authSource=%s",
		cfg.MongoDB.User,
		url.QueryEscape(cfg.Password),
		cfg.MongoDB.Host,
		cfg.MongoDB.Port,
		cfg.MongoDB.DBName,
		cfg.MongoDB.AuthMechanism,
		cfg.MongoDB.AuthDatabase,
	)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return mongoConnection, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	if client == nil {
		return mongoConnection, fmt.Errorf("failed to connect to MongoDB")
	}

	log.Info("MongoDB connected", slog.Any("mongoClient", client))

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Info("Starting MongoDB connection")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Closing MongoDB connection")
			return client.Disconnect(ctx)
		},
	})

	mongoConnection = MongoConnection{
		MongoClient: client,
		DBName:      cfg.MongoDB.DBName,
	}

	return mongoConnection, nil
}
