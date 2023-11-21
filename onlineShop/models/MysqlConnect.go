package models

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

// 初始化数据库连接
func init() {
	mysqlAdmin := beego.AppConfig.String("mysqlAdmin")
	mysqlPwd := beego.AppConfig.String("mysqlPwd")
	mysqlDb := beego.AppConfig.String("mysqlDb")
	DB, err = gorm.Open("mysql", mysqlAdmin+":"+mysqlPwd+"@/"+mysqlDb+"?charset=utf8"+
		"&parseTime=True&loc=Local")
	if err != nil {
		//待定
		/*controllers.ExecError("init mysql db failed", err, controllers.DBError, controllers.InitError)*/
	} else {
		fmt.Println("mysql connect successful")
	}
}
