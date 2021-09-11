package controllers

import (
	"beego-bbs/models"
	"beego-bbs/pkg/paginator"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
	*BaseController
}

func (t *TopicController) Index()  {
	var topics []*models.Topic
	qs := models.Topics()
	cnt, _ := models.CountObjects(qs)
	pager := paginator.NewPaginator(t.Ctx.Request, 5, cnt)
	qs = qs.Limit(5, pager.Offset()).RelatedSel()
	models.ListObjects(qs, &topics)
	t.Layout = "layouts/app.html"
	t.Data["title"] = "话题列表"
	//c.Data["class"] = "login"
	t.Data["paginator"] = pager
	t.Data["Topics"] = topics
	t.TplName = "topics/index.html"
}