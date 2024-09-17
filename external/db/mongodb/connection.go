package mongodb

import (
	"context"
	"fmt"
	"go-hdflex/external/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"log/slog"
	"net/url"
)

type Connection struct {
	Client *mongo.Client
	DbName string
}

func NewMongoDatabase(lc fx.Lifecycle, log *slog.Logger, cfg *config.Config) (Connection, error) {
	var mongoConnection Connection

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authMechanism=%s&authSource=%s",
		cfg.MongoDB.User,
		url.QueryEscape(cfg.MongoDB.Password),
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

	mongoConnection = Connection{
		Client: client,
		DbName: cfg.MongoDB.DBName,
	}

	return mongoConnection, nil
}
