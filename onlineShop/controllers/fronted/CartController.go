package fronted

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"onlineShop/models"
)

type CartController struct {
	beego.Controller
}

func (c *CartController) AddToCart() {
	ub := c.GetSession("user")
	var user models.User
	if ub != nil {
		json.Unmarshal([]byte(ub.(string)), &user)
	} else {
		user = models.GetVisitor()
		c.Ctx.WriteString("notLogin")
		return
	}
	productId, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("idErr")
		return
	}
	models.AddToCart(user.Id, productId)
	c.Ctx.WriteString("right")
}
