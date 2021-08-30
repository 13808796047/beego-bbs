package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.Layout = "layouts/app.html"
	c.Data["title"] = "首页"
	c.TplName = "home.html"
}
