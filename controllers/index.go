package controllers

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

