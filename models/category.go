package models

import "github.com/astaxie/beego/orm"

func (t *Category) Read(fields ...string) error {
	if err:= orm.NewOrm().Read(&t, fields...);err != nil {
		return err
	}
	return nil
}