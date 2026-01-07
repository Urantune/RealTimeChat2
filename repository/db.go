package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := "postgres://urantune:Seigakartisde9@localhost:5432/RealTimeChat?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	//if err := db.AutoMigrate(&models.User{}); err != nil {
	//	return err
	//}
	//
	//if err := db.AutoMigrate(&models.ChatRoom{}); err != nil {
	//	return err
	//}
	//
	//if err := db.AutoMigrate(&models.ChatMessenger{}, &models.User{}, &models.ChatRoom{}); err != nil {
	//	return err
	//}

	DB = db
	return nil
}
