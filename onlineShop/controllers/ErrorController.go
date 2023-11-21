package controllers

import (
	"github.com/astaxie/beego"
)

//集中处理报错，并储存到数据库

type ErrorController struct {
	beego.Controller
}

// 处理错误方法，返回错误类型
type errFunc func(error) string

// 将报错信息储存到数据库
func logToDB(category []string, message string, exactly string) {

}

// 依次执行处理错误的方法，并储存到数据库
func ExecError(message string, err error, errFuncS ...errFunc) {
	var errors = make([]string, len(errFuncS))
	for i, f := range errFuncS {
		errors[i] = f(err)
	}
	if err == nil {
		logToDB(errors, message, "nil")

	} else {
		logToDB(errors, message, err.Error())
	}
}
func DBError(err error) string {
	return "db"
}
func InitError(err error) string {
	return "init"
}
