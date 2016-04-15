package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"invenio/models"
)

// Operations about api
type LoginController struct {
	beego.Controller
}

// @Title Signup
// @Description Signup API
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200
// @router / [post]
func (o *LoginController) Post() {
	email := o.Input().Get("email")
	pass := o.Input().Get("pass")
	or := orm.NewOrm()
	user := models.User{
		Email:    email,
		Password: pass,
	}
	err := or.Read(&user, "email", "password")
	if err == orm.ErrNoRows {
		o.Data["json"] = models.SignUpResponse{
			Response: false,
		}
	} else {
		o.Data["json"] = models.SignUpResponse{
			Response: true,
		}
	}
	o.ServeJSON()
}
