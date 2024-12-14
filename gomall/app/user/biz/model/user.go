package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) Tablename() string {
	return "user"
}
