package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20210830_180017 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20210830_180017{}
	m.Created = "20210830_180017"

	migration.Register("User_20210830_180017", m)
}

// Run the migrations
func (m *User_20210830_180017) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("user", "innodb", "utf8mb4")
	m.PriCol("id").SetAuto(true).SetDataType("int").SetUnsigned(true)
	m.NewCol("name").SetDataType("varchar(255)").SetNullable(false)
	m.NewCol("email").SetDataType("varchar(255)").SetNullable(false)
	m.NewCol("password").SetDataType("varchar(255)").SetNullable(false)
	m.NewCol("remember_token").SetDataType("varchar(255)").SetNullable(true)
	m.NewCol("email_verified_at").SetDataType("varchar(255)").SetNullable(true)
	m.NewCol("deleted_at").SetDataType("timestamp").SetNullable(true)
	m.NewCol("created_at").SetDataType("timestamp").SetDefault("NOW()")
	m.NewCol("updated_at").SetDataType("timestamp").SetNullable(false).SetDefault("NOW()")
	m.SQL(m.GetSQL())
}

// Reverse the migrations
func (m *User_20210830_180017) Down() {
	m.SQL("DROP TABLE IF EXISTS users")

}
