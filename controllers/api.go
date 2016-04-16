package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about api
type ApiController struct {
	beego.Controller
}

// @Title Api
// @Description sentimental analysis
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200
// @router / [post]
func (o *ApiController) Post() {
	o.Data["json"] = "Hello"
	o.ServeJSON()
}
