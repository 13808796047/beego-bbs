package main

import (
	_ "beego-bbs/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
}
func main() {
	beego.Run()
}
