package controllers

import (
	"webscanner_golang/models"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type DomainScannerController struct {
	BaseController
}

var (
	domainScannerModel = new(models.DomainScanner)
)

func (this *DomainScannerController) DomainScannerMonitor() {
	//this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "scanner/domain_monitor.html"
}

func (this *DomainScannerController) DomainScannerResultList() {
	w := bson.M{}
	url := this.Input().Get("domain")
	if url != "" {
		w = bson.M{"source_domain": bson.M{"$regex": url, "$options": "$i"}}
	}
	page, _ := strconv.Atoi(this.Input().Get("page"))
	page = page - 1
	lim, _ := strconv.Atoi(this.Input().Get("limit"))
	e, countResult := domainScannerModel.CountAll(w)
	e, result := domainScannerModel.FindAll(w, page, lim)
	if !e {
		this.Data["json"] = map[string]interface{}{"code": 0, "count": 0, "data": "", "msg": ""}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "count": countResult, "data": result, "msg": "send success"}
	}
	this.ServeJSON()
}

func (this *DomainScannerController) DomainScannerResultDelete(){
	id := this.Input().Get("_id")
	if id != "" {
		v := bson.M{"_id": bson.ObjectIdHex(id)}
		e := domainScannerModel.Delete(v)
		if e {
			this.Data["json"] = map[string]interface{}{"status": success}
		} else {
			this.Data["json"] = map[string]interface{}{"status": fail}
		}
	} else {
		this.Data["json"] = map[string]interface{}{"status": fail}
	}
	this.ServeJSON()
}