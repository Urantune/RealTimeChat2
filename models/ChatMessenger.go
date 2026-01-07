package models

type ChatMessenger struct {
	ID      uint `gorm:"primaryKey"`
	Message string

	ChatRoomID uint
	ChatRoom   ChatRoom `gorm:"foreignKey:ChatRoomID"`

	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}
