package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (cfg *Config) OpenDB() (*mongo.Database, error) {
	duration, err := time.ParseDuration(cfg.Db.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	logOptions := options.
		Logger().
		SetComponentLevel(options.LogComponentCommand, options.LogLevelInfo)

	opts := options.Client().
		ApplyURI(cfg.Db.DSN).
		SetLoggerOptions(logOptions).
		SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)).
		SetMaxPoolSize(uint64(cfg.Db.MaxPoolSize)).
		SetMinPoolSize(uint64(cfg.Db.MinPoolSize)).
		SetMaxConnIdleTime(duration)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	ctx, cleanup := context.WithTimeout(context.Background(), 3*time.Second)
	defer cleanup()

	if err := client.Ping(ctx, readpref.Nearest()); err != nil {
		return nil, err
	}

	return client.Database(cfg.Db.DBName), err
}
