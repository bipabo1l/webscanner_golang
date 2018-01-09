package routers

import (
	"lazyscanner_golang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/DomainScannerMonitor", &controllers.DomainScannerController{}, "get:DomainScannerMonitor")
	beego.Router("/DomainScannerResultList", &controllers.DomainScannerController{}, "get:DomainScannerResultList")
	beego.Router("/DomainScannerResultDelete", &controllers.DomainScannerController{}, "post:DomainScannerResultDelete")
}
