package configs

import (
	"github.com/No-name16/InnoTaxi-User/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func GetConfigPostgres() (repository.Config, error) {
	config := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	validate := validator.New()
	err := validate.Struct(config)
	return config, err
}

func GetConfigMongo() (repository.ConfigMongo, error) {
	config := repository.ConfigMongo{
		Host:     viper.GetString("mongodb.host"),
		Port:     viper.GetString("mongodb.port"),
		Username: viper.GetString("mongodb.username"),
		Password: viper.GetString("mongodb.password"),
		DBName:   viper.GetString("mongodb.dbname"),
		AuthDB:   viper.GetString("mongodb.authdb"),
	}
	validate := validator.New()
	err := validate.Struct(config)
	return config, err
}
