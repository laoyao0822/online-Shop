package routers

import (
	"github.com/astaxie/beego"
	"onlineShop/controllers/fronted"
)

func init() {
	beego.Router("/index", &fronted.IndexController{}, "get:Index")
}
