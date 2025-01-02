package database

import (
	profDomain "eden/modules/profile/domain"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Ошибки для работы с базой данных
var (
	ErrDBConnectionFailed = errors.New("failed to connect to database")
	ErrDBMigrationFailed  = errors.New("failed to migrate database schema")
)

// InitGormDB подключает и инициализирует GORM для работы с базой данных.
func InitGormDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("%v: %v", ErrDBConnectionFailed, err)
	}

	// Выполнение миграции схемы базы данных
	err = db.AutoMigrate(&profDomain.Profile{}, &profDomain.Photo{})
	if err != nil {
		log.Fatalf("%v: %v", ErrDBMigrationFailed, err)
	}

	log.Println("Database connected and migrated successfully")
	return db
}
