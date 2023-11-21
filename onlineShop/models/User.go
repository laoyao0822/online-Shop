package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id       int
	Name     string
	Password string
	Email    string
}

func IsNameExited(name string) (flag bool) {
	var user User
	err := DB.Where("name=?", name).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		flag = false
		return
	} else if err == nil {
		flag = true
		return
	} else {
		flag = true
		return
	}
}
func GetVisitor() User {
	return User{
		Id:   0,
		Name: "visitor",
	}
}
