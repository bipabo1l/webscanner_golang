package routers

import (
	"webscanner_golang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/DomainScannerMonitor", &controllers.DomainScannerController{}, "get:DomainScannerMonitor")
	beego.Router("/DomainScannerResultList", &controllers.DomainScannerController{}, "get:DomainScannerResultList")
	beego.Router("/DomainScannerResultDelete", &controllers.DomainScannerController{}, "post:DomainScannerResultDelete")
	beego.Router("/IPScannerMonitor", &controllers.IPScannerController{}, "get:IPScannerMonitor")
	beego.Router("/IPScannerResultList", &controllers.IPScannerController{}, "get:IPScannerResultList")
	beego.Router("/IPScannerResultDelete", &controllers.IPScannerController{}, "post:IPScannerResultDelete")
}
