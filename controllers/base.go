package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

const (
	success   = 0
	fail      = -1
	dept_fail = -5
)

func (this *BaseController) JsonEncode(code int, error_msg string, data interface{}, count int) {
	this.Data["json"] = map[string]interface{}{"code": code, "error_msg": error_msg, "msg": error_msg, "data": data, "count": count}
	this.ServeJSON()
	this.StopRun()
}
