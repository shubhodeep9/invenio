package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"invenio/models"
	//"strconv"
	"math"
	"sort"
	"sync"
)

type RestController struct {
	beego.Controller
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func contains(s [5]string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func wellsort(a map[string]float64) string {
	var p []float64
	for _, val := range a {
		p = append(p, val)
	}
	sort.Float64s(p)
	for k, v := range a {
		if p[0] == v {
			return k
		}
	}
	return "a"
}

type Pair struct {
	Key   string
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func restsort(a map[string]float64) PairList {
	pl := make(PairList, len(a))
	i := 0
	for k, v := range a {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	return pl
}

func sailsman(lat, lng map[string]float64, ob map[string]string) [5]string {
	var sorted [5]string
	sorted[0] = "1"
	current := 0
	for j := 0; j < 4; j++ {
		fmt.Println(j)
		store := make(map[string]float64)
		for i := range ob {

			if !contains(sorted, i) {
				store[i] = Distance(lat[sorted[current]], lng[sorted[current]], lat[i], lng[i])
			}
		}
		current += 1
		sorted[current] = wellsort(store)
	}
	return sorted
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
	var resp, resp2 models.Hotels
	lat := make(map[string]float64)
	lng := make(map[string]float64)
	err := req.ToJSON(&resp)
	req = httplib.Get("https://maps.googleapis.com/maps/api/place/textsearch/json?query=tourist+places+in+Bangalore&key=AIzaSyCMipFvgdKvfJ0AGMB6gMfksOHW6kLtiq4")
	req.ToJSON(&resp2)
	if err == nil {
		for i := range resp.Results {
			if resp.Results[i].Name == ob["1"] {
				lat["1"] = resp.Results[i].Geometry.Location.Lat
				lng["1"] = resp.Results[i].Geometry.Location.Lng
				break
			}
		}
		for i := range ob {
			if i != "1" {
				for j := range resp2.Results {
					if resp2.Results[j].Name == ob[i] {
						lat[i] = resp2.Results[j].Geometry.Location.Lat
						lng[i] = resp2.Results[j].Geometry.Location.Lng
						break
					}
				}
			}
		}
		sorted := sailsman(lat, lng, ob)
		req = httplib.Get("https://maps.googleapis.com/maps/api/place/textsearch/json?query=restaurants+in+Bangalore&key=AIzaSyCMipFvgdKvfJ0AGMB6gMfksOHW6kLtiq4")
		req.ToJSON(&resp)
		var breakfast, lunch, dinner []models.RestResponse
		var wg sync.WaitGroup
		wg.Add(4)
		go func() {
			defer wg.Done()
			restu := make(map[string]models.LocStruct)
			sorted1 := make(map[string]float64)
			for i := range resp.Results {
				restu[resp.Results[i].Name] = models.LocStruct{
					Lat: resp.Results[i].Geometry.Location.Lat,
					Lng: resp.Results[i].Geometry.Location.Lng,
				}
				sorted1[resp.Results[i].Name] = Distance(lat["1"], lng["1"], restu[resp.Results[i].Name].Lat, restu[resp.Results[i].Name].Lng)

			}
			pl := restsort(sorted1)
			for i := range pl {
				for j := range resp.Results {
					if resp.Results[j].Name == pl[i].Key {
						l := models.RestResponse{
							Name:   pl[i].Key,
							Rating: resp.Results[j].Rating,
						}
						breakfast = append(breakfast, l)
					}
				}
			}
		}()
		go func() {
			defer wg.Done()
			restu := make(map[string]models.LocStruct)
			sorted1 := make(map[string]float64)
			for i := range resp.Results {
				restu[resp.Results[i].Name] = models.LocStruct{
					Lat: resp.Results[i].Geometry.Location.Lat,
					Lng: resp.Results[i].Geometry.Location.Lng,
				}
				sorted1[resp.Results[i].Name] = Distance(lat[sorted[2]], lng[sorted[2]], restu[resp.Results[i].Name].Lat, restu[resp.Results[i].Name].Lng)

			}
			pl := restsort(sorted1)
			for i := range pl {
				for j := range resp.Results {
					if resp.Results[j].Name == pl[i].Key {
						l := models.RestResponse{
							Name:   pl[i].Key,
							Rating: resp.Results[j].Rating,
						}
						lunch = append(lunch, l)
					}
				}
			}
		}()
		go func() {
			defer wg.Done()
			restu := make(map[string]models.LocStruct)
			sorted1 := make(map[string]float64)
			for i := range resp.Results {
				restu[resp.Results[i].Name] = models.LocStruct{
					Lat: resp.Results[i].Geometry.Location.Lat,
					Lng: resp.Results[i].Geometry.Location.Lng,
				}
				sorted1[resp.Results[i].Name] = Distance(lat[sorted[4]], lng[sorted[4]], restu[resp.Results[i].Name].Lat, restu[resp.Results[i].Name].Lng)

			}
			pl := restsort(sorted1)
			for i := range pl {
				for j := range resp.Results {
					if resp.Results[j].Name == pl[i].Key {
						l := models.RestResponse{
							Name:   pl[i].Key,
							Rating: resp.Results[j].Rating,
						}
						dinner = append(dinner, l)
					}
				}
			}
		}()
		var ite []models.IteStruct
		go func() {
			wg.Done()
			var ite2 models.IteStruct
			ite2 = models.IteStruct{
				Name: ob["1"],
				Type: "Hotel",
			}
			ite = append(ite, ite2)
			for i := 1; i < 5; i++ {
				ite2 = models.IteStruct{
					Name: ob[sorted[i]],
					Type: "Tourist Places",
				}
				ite = append(ite, ite2)
			}
		}()
		wg.Wait()
		o.Data["json"] = models.RestRespStruct{
			BreakFast: breakfast,
			Lunch:     lunch,
			Dinner:    dinner,
			Itenary:   ite,
		}
	}
	fmt.Println(err)
	o.ServeJSON()
}
