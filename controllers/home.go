package controllers

import (
	"beego-bbs/pkg/auth"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {


	c.Layout = "layouts/app.html"
	c.Data["title"] = "首页"
	c.Data["User"] = auth.User(c.StartSession())
	c.TplName = "home.html"
}
