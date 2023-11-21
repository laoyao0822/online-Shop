package backend

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"onlineShop/models"
	"os"
	"strconv"
)

type ProductController struct {
	beego.Controller
}

func (c *ProductController) AddProduct() {
	ub := c.GetSession("adi")
	var user models.Administrator
	if ub != nil {
		json.Unmarshal([]byte(ub.(string)), &user)
	} else {
		c.Redirect("/backend/login", 302)
	}
	c.TplName = "backend/product/add.html"
	c.Data["isSuc"] = ""
	c.Render()
}
func (c *ProductController) ToAddProduct() {
	ub := c.GetSession("adi")
	var user models.Administrator
	if ub != nil {
		json.Unmarshal([]byte(ub.(string)), &user)
	} else {
		c.Redirect("/backend/login", 302)
	}
	fmt.Println(c.Ctx.Request.Body)
	name := c.GetString("name")
	fmt.Println(c.GetString("name"))
	profile := c.GetString("profile")
	price, err := c.GetInt("price")
	fmt.Println(c.GetString("price"))
	if err != nil {
		fmt.Println(err.Error())
		c.Ctx.ResponseWriter.Write([]byte("fail"))
		return
	}
	models.DB.Transaction(func(tx *gorm.DB) error {
		product := models.Product{
			Name:    name,
			Price:   price,
			Sold:    0,
			Profile: profile,
		}
		models.DB.Create(&product)
		id := product.Id
		if id <= 0 {
			return *new(error)
		}
		_, err = os.Create("./static/img/product/Show/" + strconv.Itoa(id) + ".jpg")
		if err != nil {
			fmt.Println("create err")
			return err
		}
		_, err = os.Create("./static/img/product/Desc/" + strconv.Itoa(id) + ".jpg")
		if err != nil {
			fmt.Println("create err")
			return err
		}
		c.SaveToFile("show", "./static/img/product/Show/"+strconv.Itoa(id)+".jpg")
		c.SaveToFile("des", "./static/img/product/Desc/"+strconv.Itoa(id)+".jpg")
		models.LogAddProduct(user.Id, product)
		return nil
	})
	c.TplName = "backend/product/add.html"
	c.Data["isSuc"] = "success"
	c.Render()
}
