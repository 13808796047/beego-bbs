package models

import (
	"time"
)
type User struct {
	Id       int    `form:"_"`
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Password_confirm string `form:"password_confirm" orm:"-"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	//Topics []*Topic `orm:"reverse(many)"`
}
type Category struct {
	Id int
	Name string
	Description string `orm:"null"`
	PostCount int       `orm:"default(0)"`
	Created      time.Time `orm:"auto_now_add"`
	Updated      time.Time `orm:"auto_now;index"`
}
type Topic struct {
	Id int                  `orm:"pk"`
	Title string
	Body string
	ReplyCount      uint64 `orm:"default(0)"`
	ViewCount       uint64 `orm:"default(0)"`
	LastReplyUserId uint64 `orm:"default(0)"`
	Order           uint64 `orm:"default(0)"`
	Except          string `orm:"null"`
	Slug            string `orm:"null"`
	User         *User     `orm:"rel(fk)"`
	Category         *Category     `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}
