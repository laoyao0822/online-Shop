package models

import (
	"strconv"
	"time"
)

type LogOfAdi struct {
	adiId  int
	time   time.Time
	detail string
	method string
}

func (LogOfAdi) TableName() string {
	return "logofadi"
}
func LogAddProduct(id int, p Product) {
	DB.Create(&LogOfAdi{
		adiId:  id,
		time:   time.Now(),
		detail: strconv.Itoa(id) + "add product,product id is" + strconv.Itoa(p.Id),
		method: "addProduct",
	})
}
