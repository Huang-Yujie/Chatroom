package model

type Message struct {
	ID        uint64 `json:"message_id" gorm:"autoIncrement"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
	UserID    uint64 `json:"user_id"`
}
