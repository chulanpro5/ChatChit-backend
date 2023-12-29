package common

import (
	"test-chat/config"
	"test-chat/pkg/db/postgres"
	"test-chat/pkg/redis"
)

var common *Common

type Common struct {
	Config   *config.Config
	Database *postgres.Database
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

	redisClient, err := redis.NewRedisClient(appConfig)
	if err != nil {
		panic(err)
	}

	common = &Common{
		appConfig,
		db,
		redisClient,
	}
	return common
}

func GetCommon() *Common {
	return common
}

func (c *Common) CleanUp() {

}
