package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["invenio/controllers:ApiController"] = append(beego.GlobalControllerRouter["invenio/controllers:ApiController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["invenio/controllers:LoginController"] = append(beego.GlobalControllerRouter["invenio/controllers:LoginController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["invenio/controllers:ObjectController"] = append(beego.GlobalControllerRouter["invenio/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["invenio/controllers:ObjectController"] = append(beego.GlobalControllerRouter["invenio/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["invenio/controllers:ObjectController"] = append(beego.GlobalControllerRouter["invenio/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["invenio/controllers:ObjectController"] = append(beego.GlobalControllerRouter["invenio/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["invenio/controllers:ObjectController"] = append(beego.GlobalControllerRouter["invenio/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["invenio/controllers:SignUpController"] = append(beego.GlobalControllerRouter["invenio/controllers:SignUpController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

}
