package backend

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"onlineShop/models"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	//权限检测
	ub := c.GetSession("adi")
	var user models.Administrator
	if ub != nil {
		json.Unmarshal([]byte(ub.(string)), &user)
	} else {
		c.Redirect("/backend/login", 302)
	}
	fmt.Println(user)
	fmt.Println(user.Id, "------")
	c.TplName = "backend/index/index.html"
	c.Data["Name"] = user.Name
	c.Render()

}
