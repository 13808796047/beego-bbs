package controllers

import (
	"beego-bbs/models"
	"beego-bbs/pkg/paginator"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Show()  {

	id:= c.Ctx.Input.Param(":id")
	order:= c.GetString("order")
	var topics []*models.Topic
	category := models.Category{
		Id: cast.ToInt(id),
	}
	orm.NewOrm().Read(&category)

	query := models.Topics(order).Filter("Category",id)
	num,_ := query.Count()
	pager := paginator.NewPaginator(c.Ctx.Request, 5, num)
	query.Limit(5, pager.Offset()).RelatedSel().All(&topics)
	c.Layout = "layouts/app.html"
	c.Data["title"] = "话题列表"
	c.Data["class"] = "categories-show"
	c.Data["paginator"] = pager
	c.Data["Topics"] = topics
	c.Data["Category"] =category
	c.Data["rUrl"] = c.Ctx.Request.RequestURI
	c.TplName = "topics/index.html"
}
