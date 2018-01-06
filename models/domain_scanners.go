package models

import (
	"gopkg.in/mgo.v2/bson"
)

func (this *DomainScanner) TableName() string {
	return "domain_info"
}

type DomainScanner struct {
	Id           bson.ObjectId `bson:"id"`
	Url          string        `bson:"url"`
	Fingerprint  []string      `bson:"fingerprint"`
	CreateTime   string        `bson:"createtime"`
	Title        string        `bson:"title"`
	SourceDomain string        `bson:"sourcedomain"`
}

var (
	domainscannerModelObject = BaseModel{"", "webscanner", "domain_info"}.NewTable()
)

func (this *DomainScanner) FindAll(w bson.M, skipNum int, lim int) (bool, []map[string]interface{}) {
	var mapResult []map[string]interface{}
	var err error
	if skipNum == -1 {
		err = domainscannerModelObject.Find(w).Sort("-_id").All(&mapResult)
	} else {
		err = domainscannerModelObject.Find(w).Sort("-_id").Skip(skipNum * lim).Limit(lim).All(&mapResult)
	}
	if len(mapResult) == 0 || err != nil {
		return false, mapResult
	}
	return true, mapResult
}

func (this *DomainScanner) Update(m bson.M) bool {
	domainscannerModelObject.Update(m, bson.M{"$set": bson.M{"issendmail": 1}})
	return true
}

func (this *DomainScanner) CountAll(w bson.M) (bool, int) {
	n, err := domainscannerModelObject.Find(w).Count()
	if err != nil {
		return false, 0
	}
	return true, n
}

func (this *DomainScanner) Delete(w bson.M) bool {
	err := domainscannerModelObject.Remove(w)
	if err != nil {
		return false
	}
	return true
}

func (this *DomainScanner) InsertOne(w bson.M) bool {
	err := domainscannerModelObject.Insert(w)
	if err != nil {
		return false
	}
	return true
}
