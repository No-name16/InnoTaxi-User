package main

import (
	"context"
	"github.com/pressly/goose/v3"
	"net/http"

	"github.com/No-name16/InnoTaxi-User/configs"
	"github.com/No-name16/InnoTaxi-User/internal/handler"
	"github.com/No-name16/InnoTaxi-User/internal/repository"
	"github.com/No-name16/InnoTaxi-User/internal/service"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type dbHook struct {
	*mongo.Database
	ctx context.Context
}

func (dbHook) Levels() []log.Level {
	return []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
		log.WarnLevel,
		log.InfoLevel,
		log.DebugLevel,
	}
}

func (db dbHook) Fire(e *log.Entry) error {
	collection := db.Collection("logs")
	_, err := collection.InsertOne(db.ctx, bson.M{
		"level":   e.Level.String(),
		"created": e.Time.Format("15:04:05 3:04:05 PM"),
		"msg":     e.Message,
	})

	return err
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading enveroment variables: %s", err.Error())
	}

	ctx := context.TODO()
	configMongo, err := configs.GetConfigMongo()
	if err != nil {
		log.Fatalf("error setting mongo configs: %s", err.Error())
	}
	dblog, err2 := repository.NewMongoDB(context.Background(), configMongo)
	if err2 != nil {
		log.Fatalf("failed to initialize db for logs: %s", err.Error())
	}
	log.AddHook(dbHook{ctx: ctx, Database: dblog})
	log.SetFormatter(&log.JSONFormatter{})

	configPostgres, err := configs.GetConfigPostgres()
	if err != nil {
		log.Fatalf("error setting mongo configs: %s", err.Error())
		return
	}
	db, err := repository.NewPostgresDB(configPostgres)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Errorf("failed to set dialect to goose: %s", err.Error())
	}

	if err := goose.Up(db.DB, "internal/db"); err != nil {
		log.Errorf("failed to do migration: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	if err := http.ListenAndServe(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}
