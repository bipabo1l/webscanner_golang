package main

import (
	_ "webscanner_golang/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/layui", "/static/layui")
	beego.BConfig.WebConfig.TemplateLeft = "{{{"
	beego.BConfig.WebConfig.TemplateRight = "}}}"
	beego.Run()
}

