package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) Tablename() string {
	return "user"
}

func Create(ctx context.Context, db *gorm.DB, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func GetByEmail(ctx context.Context, db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
