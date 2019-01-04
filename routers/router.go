package routers

import (
	"github.com/astaxie/beego"
	"myAdmin/controllers"
)

func init() {
	beego.Router("/", &controllers.BaseController{},"get:LoginPage")
	beego.Router("/login", &controllers.BaseController{},"post:Login")
	beego.Router("/index", &controllers.IndexController{},"get:Index")
	beego.Router("/wel", &controllers.BaseController{},"get:Wel")

	beego.Router("/member-list", &controllers.MemberController{},"get:List")

}
