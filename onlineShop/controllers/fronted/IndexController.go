package fronted

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
	ub := c.GetSession("user")
	var user models.User
	if ub != nil {
		json.Unmarshal([]byte(ub.(string)), &user)
	} else {
		user = models.GetVisitor()
	}
	fmt.Println(user)
	fmt.Println(user.Id, "------")
	c.TplName = "fronted/index/index.html"
	c.Data["Name"] = user.Name
	c.Render()

}
