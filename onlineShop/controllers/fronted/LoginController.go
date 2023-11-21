package fronted

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"onlineShop/models"
)

type LoginAndRegister struct {
	beego.Controller
}

func (c *LoginAndRegister) Login() {
	c.TplName = "fronted/auth/login.html"
}
func (c *LoginAndRegister) TryLogin() {
	fmt.Println(c.Ctx.Request.Body)
	method := c.GetString("method")

	if method == "name" {
		password := c.GetString("password")
		name := c.GetString("name")
		user := new(models.User)
		models.DB.Where("name=?", name).Find(user)
		if password == user.Password {
			s, _ := json.Marshal(user)
			c.SetSession("user", string(s))
			c.Ctx.ResponseWriter.Write([]byte("right"))
		} else {
			c.Ctx.ResponseWriter.Write([]byte("false"))
		}
	}

}
func (c *LoginAndRegister) Register() {
	c.TplName = "fronted/auth/register.html"
}

// 判断注册是否成功
func (c *LoginAndRegister) TryRegister() {
	password := c.GetString("password")
	name := c.GetString("name")
	if models.IsNameExited(name) {
		c.Ctx.ResponseWriter.Write([]byte("repeat"))
		return
	}
	email := c.GetString("email")
	user := models.User{
		Password: password,
		Name:     name,
		Email:    email,
	}
	models.DB.Create(&user)
	fmt.Println("id:", user.Id)
	s, _ := json.Marshal(user)
	c.SetSession("user", string(s))
	c.Ctx.ResponseWriter.Write([]byte("yes"))
}
