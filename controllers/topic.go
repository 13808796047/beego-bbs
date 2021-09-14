package controllers

import (
	"beego-bbs/models"
	"beego-bbs/pkg/auth"
	"beego-bbs/pkg/paginator"
	"beego-bbs/pkg/upload"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
	"log"
)

type TopicController struct {
	beego.Controller
	*BaseController
}

func (t TopicController) UploadImage()  {
	f,h,_ := t.GetFile("upload_file")
	user := auth.User(t.StartSession())
	data := map[string]interface{}{
		"success":false,
		"msg":"上传失败",
		"file_path":"",
	}
	result,err:= upload.Save(h,"topics",cast.ToString(user.Id))
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	err = t.SaveToFile("upload_file",result)
	if err != nil {
		log.Println(err)
	}
	data = map[string]interface{}{
		"success":true,
		"msg":"上传成功",
		"file_path":result,
	}
	t.Data["json"] = data
	t.ServeJSON()
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
func (t *TopicController) Store()  {
	topic := models.Topic{}
	if err := t.ParseForm(&topic); err != nil {
		log.Println(err,"????")
		return
	}

	user := models.User{Id: 1}
	orm.NewOrm().Read(&user)
	category_id,_:=t.GetInt("category_id")

	category := models.Category{Id: category_id}
	orm.NewOrm().Read(&category)
	topic.BeforeSave(&user,&category)
	err := topic.Insert()
	if err != nil {
		log.Println(err)
	}
	t.Redirect("/topics/show/"+cast.ToString(topic.Id),302)
}
func (t *TopicController) Show() {
	id:= t.Ctx.Input.Param(":id")
	topic := models.Topic{
		Id: cast.ToInt(id),
	}
	orm.NewOrm().QueryTable("topic").Filter("id",id).RelatedSel().One(&topic)
	t.Layout = "layouts/app.html"
	t.Data["title"] = &topic.Title
	t.Data["topic"] = &topic
	t.Data["class"] = "topics-show"
	t.Data["rUrl"] = t.Ctx.Request.RequestURI
	t.TplName = "topics/show.html"
}