package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"invenio/models"
	"strconv"
)

type RestController struct {
	beego.Controller
}

// @Title Restaurants
// @Description List Restaurants
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200
// @router / [post]
func (o *RestController) Post() {
	var ob map[string]string
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	req := httplib.Get("https://maps.googleapis.com/maps/api/place/textsearch/json?query=hotels+in+Bangalore&key=AIzaSyCMipFvgdKvfJ0AGMB6gMfksOHW6kLtiq4")
	var resp models.Hotels
	latlng := make(map[string]string)
	err := req.ToJSON(&resp)
	if err == nil {
		for k := range ob {
			for i := range resp.Results {
				if resp.Results[i].Name == ob[k] {
					latlng[k] = strconv.FormatFloat(resp.Results[i].Geometry.Location.Lat, 'f', -1, 64) + "," + strconv.FormatFloat(resp.Results[i].Geometry.Location.Lng, 'f', -1, 64)

				}
			}
			fmt.Println(k)
		}
		ways := latlng["2"] + "|" + latlng["3"] + "|" + latlng["4"] + "|" + latlng["5"]
		origin := latlng["1"]
		req = httplib.Get("https://maps.googleapis.com/maps/api/directions/json?origin=" + origin + "&destination=" + origin + "&waypoints=optimize:true|" + ways + "&key=AIzaSyAmDb9Gv7rY8dWvEUbwyU0y3hQTz2eoatU")
		fmt.Println(latlng)
		var place models.Places
		err = req.ToJSON(&place)
		if err == nil {
			var result []string
			result = append(result, ob["1"])
			for j := range place.Routes[0].Legs {
				result = append(result, place.Routes[0].Legs[j].EndAddress)
			}
			o.Data["json"] = result[:len(result)-1]
		}
	}
	fmt.Println(err)
	o.ServeJSON()
}
