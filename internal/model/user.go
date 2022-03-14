package model

type User struct {
	ID uint64 `json:"user_id" gorm:"autoIncrement"`
	UserName string `json:"user_name"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	CreatedAt int `json:"created_at"`
  	UpdatedAt int `json:"updated_at"`
}