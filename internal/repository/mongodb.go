package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigMongo struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	Username string
	Password string
	DBName   string `validate:"required"`
	AuthDB   string
}

func NewMongoDB(ctx context.Context, cfg ConfigMongo) (db *mongo.Database, err error) {
	var mongoDBURL string
	var isAuth bool
	if cfg.Username == "" && cfg.Password == "" {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)
	} else {
		isAuth = true
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	}

	clientOptions := options.Client().ApplyURI(mongoDBURL)

	if isAuth {
		if cfg.AuthDB == "" {
			cfg.AuthDB = cfg.DBName
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: cfg.AuthDB,
			Username:   cfg.Username,
			Password:   cfg.Password,
		})
	}
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB: %v", err)
	}

	return client.Database(cfg.DBName), nil

}
