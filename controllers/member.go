package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myAdmin/model"
	"strconv"
	"time"
)

type MemberController struct {
	BaseController
}

func(this *MemberController)List(){
	var admin []*model.Admin
	o := orm.NewOrm().QueryTable("admin")
	if startTime,endTime:=this.GetString("start"),this.GetString("end");startTime != "" && endTime != "" {
		startTime = startTime+" 00:00:00"
		endTime = endTime+" 23:59:59"
		//时间参数
		o.Filter("created_at__between",startTime,endTime)
	}

	if userName := this.GetString("username");userName != "" {
		//用户姓名
		o.Filter("name",userName)
	}
	_,err := o.All(&admin)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"查询错误"},400)
	}

	this.Data["list"] = admin
	beego.Info(admin)


	this.TplName="member-list.html"
}

func(this *MemberController)Del(){
	adminId := this.GetString("id")
	if adminId == "" {
		this.ReturnJson(map[string]string{"message":"请传入id"},400)
	}
	id,err := strconv.Atoi(adminId)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"请传入数字"},400)
	}
	delTime := time.Now().Format("2006-01-02 15:04:05")
	num,err := orm.NewOrm().QueryTable("admin").Filter("id",id).Update(orm.Params{
		"deleted_at":delTime,
	})



	


}
