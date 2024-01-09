package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"test-chat/config"
	entity2 "test-chat/pkg/entity"
	"time"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(config *config.Config) (*Database, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", config.Database.Host, config.Database.Username, config.Database.Password, config.Database.Database, config.Database.Port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		//@TODO: Define log mode in production
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
		return nil, err
	}

	return &Database{db}, nil
}

func (db *Database) Migrate() error {
	err := db.AutoMigrate(&entity2.User{}, &entity2.Room{})
	if err != nil {
		return err
	}

	return nil
}
