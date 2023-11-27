package common

import (
	"test-chat/internal/config"
	"test-chat/internal/db/mongodb"
	"test-chat/internal/db/postgres"
	"test-chat/internal/redis"
)

var common *Common

type Common struct {
	Config   *config.Config
	Database *postgres.Database
	MongoDb  *mongodb.MongoDb
	Redis    *redis.Client
}

func NewCommon() *Common {
	appConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := postgres.NewDatabase(appConfig)
	if err != nil {
		panic(err)
	}

	mongoDb, err := mongodb.NewMongoDb(appConfig)
	if err != nil {
		panic(err)
	}

	redisClient, err := redis.NewRedisClient(appConfig)
	if err != nil {
		panic(err)
	}

	common = &Common{
		appConfig,
		db,
		mongoDb,
		redisClient,
	}
	return common
}

func GetCommon() *Common {
	return common
}

func (c *Common) CleanUp() {

}
