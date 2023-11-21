package backend

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"onlineShop/models"
	"strconv"
)

type LoginAndRegister struct {
	beego.Controller
}

func (c *LoginAndRegister) Login() {
	c.TplName = "backend/auth/login.html"
}
func (c *LoginAndRegister) TryLogin() {
	fmt.Println(c.Ctx.Request.Body)
	password := c.GetString("password")
	name := c.GetString("name")
	adi := new(models.Administrator)
	models.DB.Where("name=?", name).Find(&adi)
	if password == adi.Password {
		s, _ := json.Marshal(adi)
		c.SetSession("adi", string(s))
		c.Ctx.ResponseWriter.Write([]byte("right"))
	} else {
		c.Ctx.ResponseWriter.Write([]byte("false"))
	}
}
func (c *LoginAndRegister) Register() {
	//权限检测
	ub := c.GetSession("adi")
	var user models.Administrator
	if ub != nil {
		json.Unmarshal([]byte(ub.(string)), &user)
		if !user.IsSuper {
			c.Redirect("/backend/login", 302)
		}
	} else {
		c.Redirect("/backend/login", 302)
	}

	c.TplName = "backend/auth/AddAdministrator.html"
}
func (c *LoginAndRegister) TryRegister() {
	//权限检测
	ub := c.GetSession("adi").(string)
	var user models.Administrator
	if ub != "" {
		json.Unmarshal([]byte(ub), &user)
		if !user.IsSuper {
			c.Redirect("/backend/login", 302)
		}
	} else {
		c.Redirect("/backend/login", 302)
	}

	password := c.GetString("password")
	name := c.GetString("name")
	if models.IsNameExitedOfAdm(name) {
		c.Ctx.ResponseWriter.Write([]byte("repeat"))
		return
	} else {

	}
	adi := models.Administrator{
		Password: password,
		Name:     name,
		IsSuper:  false,
	}
	models.DB.Create(&adi)
	fmt.Println("id:", adi.Id)
	c.Ctx.ResponseWriter.Write([]byte(strconv.Itoa(adi.Id)))
}
