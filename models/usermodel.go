package models

import "time"

type User struct {
	ID       uint32
	Name     string `gorm:"type:varchar(255);not null"`
	Username string `gorm:"type:varchar(255);not null, unique"`
	Email    string	`gorm:"type:varchar(255);not null"`
	Password string	`gorm:"type:varchar(255);not null"`
	Gender   string	`gorm:"type:varchar(255);not null"`
	Role     string	`gorm:"type:varchar(255);not null"`
	Course   string	`gorm:"type:varchar(255);not null"`
	Image    string	`gorm:"type:varchar(255)"`
	About    string	`gorm:"type:varchar(255)"`
	Joindate time.Time
}


