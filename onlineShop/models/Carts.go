package models

import (
	"onlineShop/common"
)

type Cart struct {
	UserId    int
	ProductId int
	AddTime   int
	Amount    int
}

// 向购物车中添加物品，存在则数量加一
func AddToCart(userId int, productId int) {
	cart := Cart{
		UserId:    userId,
		ProductId: productId,
		Amount:    0,
	}
	DB.Where("user_id=? and product_id=?", userId, productId).Find(&cart)
	addTime := int(common.GetUnix())
	amount := cart.Amount + 1
	if cart.Amount == 0 {
		DB.Save(&cart)
	} else {
		DB.Model(&cart).Where("user_id=? and product_id=?", userId, productId).Update(map[string]interface{}{"amount": amount, "add_time": addTime})
	}
}
