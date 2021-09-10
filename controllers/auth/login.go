package auth

import (
	"beego-bbs/pkg/auth"
	"errors"
	"github.com/astaxie/beego"
	"net/http"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) ShowLoginForm() {
	c.Layout = "layouts/app.html"
	c.Data["title"] = "登录"
	c.Data["class"] = "login"
	c.TplName = "auth/login.html"
}
func (c *LoginController) Login() {
	email := c.GetString("email")
	password := c.GetString("password")
	session := c.StartSession()
	if err := auth.Attempt( email, password,session); err == nil {
		c.Redirect("/",http.StatusFound)
	}else {
		c.Layout = "layouts/app.html"
		c.Data["title"] = "登录"
		c.Data["class"] = "login"
		c.Data["Errors"]= errors.New("账号或密码错误")
		c.TplName = "auth/login.html"
	}
}
func (c *LoginController) Logout() {
	session := c.StartSession()
	auth.Logout(session)
	c.Redirect("/",http.StatusFound)
}