package routers

import (
	"beego-bbs/controllers"
	"beego-bbs/controllers/auth"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/register", &auth.RegisterController{}, "get:ShowRegistrationForm")
	beego.Router("/register", &auth.RegisterController{}, "post:Register")
}
