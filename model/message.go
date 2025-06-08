package model

type Message struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}
