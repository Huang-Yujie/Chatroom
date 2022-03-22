package model

type Message struct {
	ID        uint64 `json:"message_id" gorm:"autoIncrement"`
	Content   string `json:"message_content"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
	*User     `json:"user"`
}
