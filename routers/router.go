package routers

import (
	"beego-bbs/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.HomeController{}, "get:Index")
}
