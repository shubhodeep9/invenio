package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"invenio/models"
)

// Operations about api
type ShowController struct {
	beego.Controller
}

// @Title Show Reviews
// @Description Signup API
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200
// @router / [post]
func (o *ShowController) Post() {
	or := orm.NewOrm()
	var uploads []models.Upload
	if _, err := or.QueryTable("upload").All(&uploads); err == nil {
		o.Data["json"] = uploads
	}
	o.ServeJSON()
}
