// @APIVersion 0.0.1
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"invenio/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/api",
			beego.NSInclude(
				&controllers.ApiController{},
			),
		),
		beego.NSNamespace("signup",
			beego.NSInclude(
				&controllers.SignUpController{},
			),
		),
		beego.NSNamespace("login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("showrev",
			beego.NSInclude(
				&controllers.ShowController{},
			),
		),
		beego.NSNamespace("hotels",
			beego.NSInclude(
				&controllers.HotelController{},
			),
		),
		beego.NSNamespace("restaurants",
			beego.NSInclude(
				&controllers.RestController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
