package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Acct     string    `gorm:"primary_key" json:"acct"`
	Pwd      string    `gorm:"sizr=100;not null;" json:"pwd"`
	Fullname string    `gorm:"size=255;not null;" json:"fullname"`
	CreateAt time.Time `gorm:"CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt time.Time `gorm:"CURRENT_TIMESTAMP" json:"update_at"`
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}
