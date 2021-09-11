package models

import "github.com/astaxie/beego/orm"

// 新增
func (t *Topic) Insert() error {
	if _,err := orm.NewOrm().Insert(t);err !=nil{
		return err
	}
	return nil
}
func (t *Topic) Read(fields ...string) error {
	if err:= orm.NewOrm().Read(t, fields...);err != nil {
		return err
	}
	return nil
}
func (t *Topic) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}
func Topics(order string) orm.QuerySeter {
	query := orm.NewOrm().QueryTable("topic")
	switch order {
	case "recent":
		query.OrderBy("-CreatedAt")
		break
	default:
		query.OrderBy("-UpdatedAt")
	}
	return query
}