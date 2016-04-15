package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"invenio/models"
)

// Operations about api
type SignUpController struct {
	beego.Controller
}

// @Title Signup
// @Description Signup API
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200
// @router / [post]
func (o *SignUpController) Post() {
	email := o.Input().Get("email")
	pass := o.Input().Get("pass")
	or := orm.NewOrm()
	user := models.User{
		Email:    email,
		Password: pass,
	}
	if created, _, err := or.ReadOrCreate(&user, "email", "password"); err == nil {
		o.Data["json"] = models.SignUpResponse{
			Response: created,
		}
	}
	o.ServeJSON()
}
