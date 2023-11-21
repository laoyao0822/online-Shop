package models

import "github.com/jinzhu/gorm"

type Administrator struct {
	IsSuper  bool `column:"isSuper""`
	Name     string
	Id       int
	Password string
}

func IsNameExitedOfAdm(name string) (flag bool) {
	var adi Administrator
	err := DB.Where("name=?", name).First(&adi).Error
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
