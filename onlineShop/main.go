package main

import (
	"github.com/astaxie/beego"
	_ "onlineShop/routers"
)

func main() {
	beego.SetStaticPath("/productImg", "static/img/product")
	beego.Run()
}
