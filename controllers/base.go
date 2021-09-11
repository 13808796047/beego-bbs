package controllers

import (
	"beego-bbs/pkg/paginator"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}
func (b *BaseController) SetPaginator(per int, nums int64) *paginator.Paginator {
	p := paginator.NewPaginator(b.Ctx.Request, per, nums)
	b.Data["paginator"] = p
	return p
}