package models

import "time"

type User struct {
	ID         int       `json:"id"`
	Fullname   string    `json:"fullname" gorm:"type: varchar(255)"`
	Email      string    `json:"email" gorm:"type: varchar(255)"`
	Password   string    `json:"password" gorm:"type: varchar(255)"`
	Username   string    `json:"username" gorm:"type: varchar(255)"`
	ListAsRole string    `json:"list_as_role" gorm:"type: varchar(225)"`
	Address    string    `json:"address" gorm:"type: varchar(225)"`
	Gender     string    `json:"gender" gorm:"type: varchar(225)"`
	Phone      string    `json:"phone" gorm:"type: varchar(225)"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID         int    `json:"id"`
	Fullname   string `json:"fullname" `
	Email      string `json:"email" `
	Password   string `json:"password" `
	Username   string `json:"username" `
	ListAsRole string `json:"list_as_role" `
	Address    string `json:"address" `
	Gender     string `json:"gender" `
	Phone      string `json:"phone" `
	Image      string `json:"image"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
