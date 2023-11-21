package routers

import (
	"github.com/astaxie/beego"
	"onlineShop/controllers/backend"
)

func init() {
	beego.Router("/backend/login", &backend.LoginAndRegister{}, "get:Login")
	beego.Router("/backend/login/try", &backend.LoginAndRegister{}, "get:TryLogin")
	beego.Router("/backend/register", &backend.LoginAndRegister{}, "get:Register")
	beego.Router("/backend/register/try", &backend.LoginAndRegister{}, "post:TryRegister")
	beego.Router("/backend/index", &backend.IndexController{}, "get:Index")
	beego.Router("/backend", &backend.IndexController{}, "get:Index")

	//product
	beego.Router("/backend/product/add", &backend.ProductController{}, "post:ToAddProduct")
	beego.Router("/backend/product/add", &backend.ProductController{}, "get:AddProduct")
}
