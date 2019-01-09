package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myAdmin/model"
	"strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
	isLogin bool
}
type MyPage struct {
	Count int64	//总条数
	CountPage int	//总页数
	Limit int //每页几条
	NowPage int //当前页
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

func(this *BaseController)Wel(){
	this.TplName = "welcome.html"
}


//分页
func(this *BaseController)GetPage(o orm.QuerySeter)(orm.QuerySeter,*MyPage,error){
	var myPage MyPage
	myPage.Limit = 10
	myPage.NowPage = 1

	var err3  error
	myPage.Count,err3 = o.Count()
	if err3!= nil {
		return nil,&myPage, err3
	}

	//总页数
	myPage.CountPage = int(myPage.Count)/myPage.Limit
	if m := int(myPage.Count)%myPage.Limit;m>0 {
		myPage.CountPage++
	}

	if limitStr := this.GetString("limit");limitStr !="" {
		var err2 error
		myPage.Limit,err2 = strconv.Atoi(limitStr)
		if err2 != nil {
			return nil,&myPage,err2
		}
	}
	if pageStr := this.GetString("page") ;pageStr != ""{
		var err error
		myPage.NowPage,err = strconv.Atoi(pageStr)
		if err != nil {
			return nil,&myPage ,err
		}
	}

	return o.Limit(myPage.Limit,(myPage.NowPage-1)*myPage.Limit),&myPage,nil
}


func MakeMd5(str string)string{
	has := md5.Sum([]byte(str+"yan"))
	return fmt.Sprintf("%x",has)
}