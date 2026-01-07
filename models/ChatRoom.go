package models

type ChatRoom struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
