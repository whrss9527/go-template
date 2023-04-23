package model

type User struct {
	Id string `json:"id" gorm:"primaryKey;column:id"`
}
