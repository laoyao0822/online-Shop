package routers

import (
	"github.com/astaxie/beego"
	"onlineShop/controllers/fronted"
)

func init() {
	beego.Router("/login", &fronted.LoginAndRegister{}, "get:Login")
	beego.Router("/login/try", &fronted.LoginAndRegister{}, "get:TryLogin")
	beego.Router("/login/register", &fronted.LoginAndRegister{}, "get:Register")
	beego.Router("/login/register/try", &fronted.LoginAndRegister{}, "post:TryRegister")
	beego.Router("/product/?:id", &fronted.ProductController{}, "get:GetProduct")
	beego.Router("/product/addToCart", &fronted.CartController{}, "get:AddToCart")
}
