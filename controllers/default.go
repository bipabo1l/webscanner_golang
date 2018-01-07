package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	beego.SetStaticPath("/layui", "/static/layui")
	beego.BConfig.WebConfig.TemplateLeft = "{{{"
	beego.BConfig.WebConfig.TemplateRight = "}}}"
	this.TplName = "index.html"
}
