package fronted

import (
	"github.com/astaxie/beego"
	"onlineShop/models"
	"strconv"
)

type ProductController struct {
	beego.Controller
}

func (c *ProductController) GetProduct() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Redirect("/index", 302)
	}
	p := models.GetProductById(id)
	if p.Name == "" {
		c.Redirect("/index", 302)
	}
	c.TplName = "fronted/product/product.html"
	c.Data["Name"] = p.Name
	c.Data["Profile"] = p.Profile
	c.Render()
}
