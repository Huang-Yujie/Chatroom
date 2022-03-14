package model

type Message struct {
	ID uint64 `json:"message_id" gorm:"autoIncrement"`
	CreatedAt int `json:"created_at"`
  	UpdatedAt int `json:"updated_at"`
	UserID uint64 `json:"user_id"`
}