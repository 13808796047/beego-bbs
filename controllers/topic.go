package controllers

import (
	"beego-bbs/models"
	"beego-bbs/pkg/paginator"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TopicController struct {
	beego.Controller
	*BaseController
}

func (t *TopicController) Index()  {
	var topics []*models.Topic
	order:= t.GetString("order")
	qs := models.Topics(order)
	cnt, _ := models.CountObjects(qs)
	pager := paginator.NewPaginator(t.Ctx.Request, 5, cnt)
	qs = qs.Limit(5, pager.Offset()).RelatedSel()
	models.ListObjects(qs, &topics)
	t.Layout = "layouts/app.html"
	t.Data["title"] = "话题列表"
	t.Data["class"] = "topics-index"
	t.Data["paginator"] = pager
	t.Data["Topics"] = topics
	t.Data["rUrl"] = t.Ctx.Request.RequestURI
	t.TplName = "topics/index.html"
}
func (t *TopicController)Create()  {
	var categories []*models.Category
	orm.NewOrm().QueryTable("category").All(&categories)
	t.Layout = "layouts/app.html"
	t.LayoutSections = make(map[string]string)
	t.LayoutSections["Styles"] = "topics/_styles.html"
	t.LayoutSections["Scripts"] = "topics/_scripts.html"
	t.Data["title"] = "话题创建"
	t.Data["class"] = "topics-create"
	t.Data["Categories"] = categories
	t.Data["rUrl"] = t.Ctx.Request.RequestURI
	t.TplName = "topics/create_and_edit.html"
}
func (t TopicController) Store()  {
	
}