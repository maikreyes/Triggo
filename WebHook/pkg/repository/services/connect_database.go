package services

import (
	"log"
	"triggo/pkg/repository/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (s *Services) ConnectDatabase() (*gorm.DB, error) {

	if s.DB != nil {
		return s.DB, nil
	}

	db, err := gorm.Open(postgres.Open(s.Config.DSN), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.RepositoryWebhook{})

	if err != nil {
		return nil, err
	}

	s.DB = db
	log.Println("Base de datos conectada y migrada exitosamente")

	return db, nil
}
