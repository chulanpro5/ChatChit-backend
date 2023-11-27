package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"test-chat/internal/config"
	"test-chat/internal/entity"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(config *config.Config) (*Database, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", config.Database.Host, config.Database.Username, config.Database.Password, config.Database.Database, config.Database.Port)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		//@TODO: Define log mode in production
		//Logger: log.NewGormLogger(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
		return nil, err
	}

	return &Database{db}, nil
}

func (db *Database) Migrate() error {
	err := db.AutoMigrate(&entity.User{}, &entity.Room{})
	if err != nil {
		return err
	}

	return nil
}
