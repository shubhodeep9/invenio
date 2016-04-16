package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"invenio/models"
)

// Operations about api
type HotelController struct {
	beego.Controller
}

// @Title Signup
// @Description Signup API
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200
// @router / [post]
func (o *HotelController) Post() {
	req := httplib.Get("https://maps.googleapis.com/maps/api/place/textsearch/json?query=hotels+in+Bangalore&key=AIzaSyAmDb9Gv7rY8dWvEUbwyU0y3hQTz2eoatU")
	var resp models.Hotels
	if req.ToJSON(&resp) == nil {
		for i := range resp.Results {
			resp.Results[i].Rating *= 20
		}
		o.Data["json"] = resp
	} else {
		o.Data["json"] = models.SignUpResponse{
			Response: false,
		}
	}
	o.ServeJSON()
}
