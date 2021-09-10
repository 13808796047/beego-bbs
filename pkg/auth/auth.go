package auth

import (
	"beego-bbs/models"
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	"github.com/spf13/cast"
)

func _getUID(store session.Store )interface{}{
	_uid := store.Get("uid")
	return _uid
}

// User 获取登录用户信息
func User(store session.Store) *models.User{
	uid := _getUID(store)

		o := orm.NewOrm()
		user := models.User{Id: cast.ToInt(uid)}
		err := o.Read(&user)
		if err == nil {
			return &user
		}

	return &models.User{}
}

// Attempt 登录
func Attempt(email, password string,store session.Store) error {
	// 1.根据Email获取用户
	var user *models.User
	_user,err := user.GetByEmail(email)
	if err != nil {
		return errors.New("账号不存在或错误")
	}
	// 2 匹配密码
	if !_user.ComparePassword(password) {

		return errors.New("密码不存在或错误")
	}

	store.Set("uid",_user.Id)
	return nil
}

// Login登录用户
func Login(user *models.User,store session.Store)  {
	store.Set("uid",user.Id)
}
func Logout(store session.Store)  {
	store.Delete("uid")
}
func Check(store session.Store) bool{
	return _getUID(store)==nil
}