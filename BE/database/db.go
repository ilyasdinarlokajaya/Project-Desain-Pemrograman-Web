package database

import (
	"dpw-latihan/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Datauser{})
	db.AutoMigrate(&model.Databanner{})
	db.AutoMigrate(&model.Datanotification{})
	db.AutoMigrate(&model.Dataproduct{})
	db.AutoMigrate(&model.Dataplannogram{})
	db.AutoMigrate(&model.Datavoucher{})
	db.AutoMigrate(&model.Dataarticle{})
	db.AutoMigrate(&model.Datadeliverymethod{})

	return db
}
