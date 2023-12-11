package modules

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OrmModule(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}
