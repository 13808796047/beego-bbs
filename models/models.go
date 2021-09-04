package models

import (
	"beego-bbs/pkg/password"
	"github.com/astaxie/beego/orm"
)
type User struct {
	Id       int    `form:"_"`
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Password_confirm string `form:"password_confirm" orm:"-"`
}
func(u *User) BeforCreate()  {
	u.Password = password.Hash(u.Password)
	return
}
func init() {
	orm.RegisterModel(&User{})
}
