package auth

import (
	"beego-bbs/models"
	"beego-bbs/requests"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/validation"
	"log"
)

type RegisterController struct {
	beego.Controller
}

var cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 90
	cpt.StdHeight = 35
}

func (c *RegisterController) ShowRegistrationForm() {
	c.Layout = "layouts/app.html"
	c.Data["title"] = "注册"
	c.Data["class"] = "register"
	c.TplName = "auth/register.html"
}

func (c *RegisterController) Register() {
	user := models.User{}
	if err := c.ParseForm(&user); err != nil {
		log.Println(err,"????")
		return
	}
	// 验证码验证
	valid := validation.Validation{}
	if !cpt.VerifyReq(c.Ctx.Request) {
		valid.SetError("captcha","验证码错误")
	}
	errors := requests.ValidateRegistrationForm(&valid,&user)

	if len(errors)>0 {
		log.Println("errr")
		c.Layout = "layouts/app.html"
		c.Data["title"] = "注册"
		c.Data["class"] = "register"
		c.Data["errors"] = errors
		c.TplName = "auth/register.html"
	}else{
		o := orm.NewOrm()
		user.BeforCreate()
		id, err := o.Insert(&user)
		if err != nil {
			log.Println(err)
		}
		if id>0 {
			c.Redirect("/", 302)
		}
	}
}
