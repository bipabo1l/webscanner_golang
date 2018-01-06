package models

import (
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
	"fmt"
)

type BaseModel struct {
	DbSource  string
	DbName    string
	TableName string
}

func (this BaseModel) NewTable() (*mgo.Collection) {
	if this.DbSource == "" {
		this.DbSource = "default"
	}
	s, err := mgo.Dial(beego.AppConfig.String(fmt.Sprintf("mongodb_%s::url", this.DbSource)))
	if err != nil {
		panic(err)
	}
	session := s
	session.SetMode(mgo.Monotonic, true)
	session.Refresh()
	return session.DB(this.DbName).C(this.TableName)
}
