package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myAdmin/model"
	"strings"
)

type BaseController struct {
	beego.Controller
	isLogin bool
}
func init(){

}
func(this *BaseController)LoginPage(){
	this.TplName = "login.html"
}

func(this *BaseController)Login(){
	pwd := this.GetString("pwd")
	user := this.GetString("user")
	if user == "" {
		this.ReturnJson(map[string]string{"message":"用户名不存在"},400)
	}
	if pwd == "" {
		this.ReturnJson(map[string]string{"message":"密码不存在"},400)
	}

	//md5加密 pwd

	has := md5.Sum([]byte(pwd+"yan"))
	pwdMB := fmt.Sprintf("%x",has)
	beego.Info(pwdMB)
	//查询user表
	var admin model.Admin
	err := orm.NewOrm().QueryTable("admin").Filter("name",user).One(&admin)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"查无此用户"},400)
	}
	beego.Info(admin)
	if admin.Password == pwdMB {
		this.SetSession("user_name",user)
		this.SetSession("user_id",admin.Id)
		this.ReturnJson(map[string]string{"message":"登录成功"},200)
	}
	this.ReturnJson(map[string]string{"message":"密码错误"},400)
}

//自己的重定向
func(this *BaseController)MyRedirect(url string){
	this.Redirect(url,302)
	this.StopRun()
}
//获得ip
func(this *BaseController)GetClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr,":")
	return s[0]
}
//自己的return json数据
func(this *BaseController)ReturnJson(data interface{},status int){
	this.Ctx.Output.Status = status
	this.Ctx.Output.JSON(data,true,false)
}