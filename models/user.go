package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *User) TableName() string {
	return "user"
}

//UploadTable

type Upload struct {
	Id          int    `json:"id"`
	ImgUrl      string `json:"img_url"`
	AgeCategory int    `json:"ageCategory"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Smile       int    `json:"smile"`
	Couple      int    `json:"couple"`
	Location    string `json:"location"`
}

func (b *Upload) TableName() string {
	return "upload"
}

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/data.db")
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Upload))
}
