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
	var req *httplib.BeegoHTTPRequest
	if o.Input().Get("i") == "1" {
		req = httplib.Get("https://maps.googleapis.com/maps/api/place/textsearch/json?query=tourist+places+in+Bangalore&key=AIzaSyCMipFvgdKvfJ0AGMB6gMfksOHW6kLtiq4")
	} else {
		req = httplib.Get("https://maps.googleapis.com/maps/api/place/textsearch/json?query=hotels+in+Bangalore&key=AIzaSyCMipFvgdKvfJ0AGMB6gMfksOHW6kLtiq4")
	}
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
