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
	beego.Router("/member-list2", &controllers.MemberController{},"get:List2")
	beego.Router("/member-list3", &controllers.MemberController{},"get:List3")
	beego.Router("/member-add", &controllers.MemberController{},"get:Add")//跳页面
	beego.Router("/member-add-admin", &controllers.MemberController{},"post:AddAdmin")//跳页面
	beego.Router("/member-list2-data", &controllers.MemberController{},"get:ListData")
	beego.Router("/member-del", &controllers.MemberController{},"delete:Del")

}
