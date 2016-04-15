package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about api
type ApiController struct {
	beego.Controller
}

// @Title Get
// @Description sentimental analysis
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200
// @router / [get]
func (o *ApiController) Get() {
	o.Data["json"] = "Hello"
	o.ServeJSON()
}
