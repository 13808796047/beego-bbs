package routers

import (
	"beego-bbs/controllers"
	"beego-bbs/controllers/auth"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/register", &auth.RegisterController{}, "get:ShowRegistrationForm")
	beego.Router("/register", &auth.RegisterController{}, "post:Register")
	beego.Router("/login", &auth.LoginController{}, "get:ShowLoginForm")
	beego.Router("/login", &auth.LoginController{}, "post:Login")
	beego.Router("/logout", &auth.LoginController{}, "post:Logout")
	beego.Router("/topics", &controllers.TopicController{},"get:Index")
	beego.Router("/categories/show/:id([0-9]+)",&controllers.CategoryController{},"get:Show")
}
