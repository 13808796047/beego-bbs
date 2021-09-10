package models

import (
	"beego-bbs/pkg/password"
	"github.com/astaxie/beego/orm"
)

func (u *User)GetByEmail(email string) (*User,error) {
	o := orm.NewOrm()
	user := &User{Email: email}
	err := o.Read(user,"Email")
	if err != nil {
		return nil,err
	}
	return user,nil
}
func (u *User) ComparePassword(_password string) bool  {
	return password.CheckHash(_password,u.Password)
}
func(u *User) BeforCreate()  {
	u.Password = password.Hash(u.Password)
}
