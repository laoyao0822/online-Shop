package models

type Product struct {
	Id      int
	Name    string
	Price   int
	Sold    int
	Profile string
}

func GetProductById(id int) (p Product) {
	DB.Where("id=?", id).Find(&p)
	return p
}
