package main

import (
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
	//user_id := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//cate_id := []int{1, 2, 3, 4}
	//var data []models.Topic
	//for i := 0; i < 100; i++ {
	//	user := models.Topic{
	//		Title:      faker.Lorem().Sentence(100),
	//		Body:        faker.Lorem().Paragraph(200),
	//		Except:      faker.Lorem().Paragraph(100),
	//	}
	//	data = append(data, user)
	//
	//}
	//
	//orm.NewOrm().InsertMulti(1, data)
	orm.RunSyncdb("default", false, true)
}
func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	// "127.0.0.1:6379,1000,$password"
	//其中127.0.0.1:6379为ip和端口，1000为连接池，最后一个为redis密码
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:63790"
	//设置日志输出的目录
	//beego.SetLogger("file", `{"filename":"logs/beego.log"}`)
	//beego.AddFuncMap("Session",helpers.GetSessionUser)
	beego.Run()
}
