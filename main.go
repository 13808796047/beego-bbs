package main

import (
	"beego-bbs/helpers"
	"beego-bbs/models"
	_ "beego-bbs/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			beego.AppConfig.String("mysqluser"),
			beego.AppConfig.String("mysqlpass"),
			beego.AppConfig.String("mysqlurls"),
			beego.AppConfig.String("mysqlport"),
			beego.AppConfig.String("mysqldb"),
		))
	orm.RegisterModel(&models.User{},&models.Topic{},&models.Category{})
	//var categories = []models.Category{{
	//	Name:        "分享",
	//	Description: "分享创造，分享发现",
	//},
	//	{
	//		Name:        "教程",
	//		Description: "开发技巧、推荐扩展包等",
	//	},
	//	{
	//		Name:        "问答",
	//		Description: "请保持友善，互帮互助",
	//	},
	//	{
	//		Name:        "公告",
	//		Description: "站点公告",
	//	}}
	//// create table
	//orm.NewOrm().InsertMulti(1, categories)

	//for i := 0; i < 14; i++ {
	//	user := models.User{
	//
	//		Name: faker.Name().Name(),
	//		Email:faker.Internet().Email(),
	//		Password:"$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi",
	//		Avatar: faker.Avatar().String(),
	//	}
	//
	//
	//	orm.NewOrm().Insert(&user)
	//}
	//user_id := []int{ 3, 4, 5, 6, 7, 8, 9, 10,11,12,13,14}
	//cate_id := []int{1, 2, 3, 4}
	//var users []*models.User
	//orm.NewOrm().QueryTable("user").All(&users)
	//var categories []*models.Category
	//orm.NewOrm().QueryTable("category").All(&categories)
	//for i := 0; i < 100; i++ {
	//
	//
	//	topic := models.Topic{
	//		Title:      faker.Lorem().String(),
	//		Body:        faker.Lorem().Sentence(10),
	//		Except:      faker.Lorem().Sentence(10),
	//	}
	//	user := models.User{Id: faker.RandomInt(3,14)}
	//
	//	orm.NewOrm().Read(&user)
	//	category := models.Category{Id: faker.RandomInt(1,4)}
	//	orm.NewOrm().Read(&category)
	//	topic.User = &user
	//	topic.Category = &category
	//	//log.Println(topic)
	//	 orm.NewOrm().Insert(&topic)
	//
	//}
	orm.RunSyncdb("default", false, true)
}
func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	// "127.0.0.1:6379,1000,$password"
	//其中127.0.0.1:6379为ip和端口，1000为连接池，最后一个为redis密码
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:63790"
	//设置日志输出的目录
	//beego.SetLogger("file", `{"filename":"logs/beego.log"}`)

	//beego.AddFuncMap("active_class",controllers.CategoryController.ActiveClass)
	//beego.AddFuncMap("topic_nav_active",controllers.CategoryController.TopicNavActive)
	beego.AddFuncMap("urlContain", helpers.UrlContain)
	beego.Run()
}
