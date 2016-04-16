package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"invenio/models"
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
	img := o.Input().Get("img")
	req := httplib.Get("http://apius.faceplusplus.com/v2/detection/detect?api_key=e2707513a30c55f950583457e8845ec1&api_secret=9cWd6oDOtFMmqhGT7mwPKphefakx52tI&url=" + img)
	var response models.ApiStruct
	if req.ToJSON(&response) == nil {
		var smile float64 = 0
		var age float64 = 0
		var people int = len(response.Face)
		for i := range response.Face {
			smile = smile + response.Face[i].Attribute.Smiling.Value
			age = age + float64(response.Face[i].Attribute.Age.Value)
		}
		smile = smile / float64(people)
		age = age / float64(people)
		var ageCat int
		if age >= 0 && age <= 12 {
			ageCat = 0
		} else if age >= 13 && age <= 17 {
			ageCat = 1
		} else if age >= 18 && age <= 25 {
			ageCat = 2
		} else if age >= 26 && age <= 45 {
			ageCat = 3
		} else {
			ageCat = 4
		}
		var couple int = 0
		if people == 2 {
			if (response.Face[0].Attribute.Gender.Value == "Male" && response.Face[1].Attribute.Gender.Value == "Female") || (response.Face[1].Attribute.Gender.Value == "Male" && response.Face[0].Attribute.Gender.Value == "Female") {
				couple = 1
			}
		}
		lat, long := o.Input().Get("latitude"), o.Input().Get("longitude")
		req = httplib.Get("http://maps.googleapis.com/maps/api/geocode/json?latlng=" + lat + "," + long + "&sensor=true")
		var geores models.GeoCode
		if req.ToJSON(&geores) == nil {
			location := geores.Result[0].Formatted
			or := orm.NewOrm()
			dbins := models.Upload{
				ImgUrl:      img,
				AgeCategory: ageCat,
				Latitude:    lat,
				Longitude:   long,
				Smile:       int(smile),
				Couple:      couple,
				Location:    location,
			}
			fmt.Println(or.Insert(&dbins))
			o.Data["json"] = models.SignUpResponse{
				Response: true,
			}
		} else {
			o.Data["json"] = models.SignUpResponse{
				Response: false,
			}
		}
	} else {
		o.Data["json"] = models.SignUpResponse{
			Response: false,
		}
	}
	o.ServeJSON()
}
