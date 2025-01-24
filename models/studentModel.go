package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Student struct {
	ID    int    `orm:"auto"`
	Name  string `orm:"size(25)"`
	Email string `orm:"size(50);unique"`
}

func init() {
	orm.RegisterModel(new(Student))
}
