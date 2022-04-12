package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string      `gorm:"not null" json:"name" form:"name"`
	City      string      `gorm:"not null" json:"city" form:"city"`
	UrlImage  string      `gorm:"not null" json:"url_image" form:"url_image"`
	Email     string      `gorm:"unique;not null" json:"email" form:"email"`
	Password  string      `gorm:"not null" json:"password" form:"password"`
	Event     []Event     `gorm:"foreignKey:UserId;references:ID"`
	Attendees []Attendees `gorm:"foreignKey:UserId;references:ID"`
	Comment   []Comment   `gorm:"foreignKey:UserId;references:ID"`
}
