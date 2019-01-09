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

/*
		用户列表  ajax渲染
 */
func(this *MemberController)List(){
	var admin []*model.Admin
	o := orm.NewOrm().QueryTable("admin")
	if startTime,endTime:=this.GetString("start"),this.GetString("end");startTime != "" && endTime != "" {
		startTime = startTime+" 00:00:00"
		endTime = endTime+" 23:59:59"
		//时间参数
		o = o.Filter("created_at__between",startTime,endTime)
	}

	if userName := this.GetString("username");userName != "" {
		//用户姓名
		o =o.Filter("name",userName)
	}
	o =o.Filter("deleted_at__isnull",true)
	//count,err3 := o.Count()
	beego.Info("member-getPage之前")
	o,myPage,err3:=this.GetPage(o)
	if err3 != nil {
		this.ReturnJson(map[string]string{"message":"分页错误"},400)
	}

	_,err :=o.All(&admin)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"查询错误"},400)
	}
	////每页几条
	//limit := 10
	////总页数
	//countPage := int(count)/limit
	//if m := int(count)%limit;m>0 {
	//	countPage++
	//}
	this.Data["count_page"] = myPage.CountPage
	this.Data["list"] = admin
	this.Data["count"] = myPage.Count
	this.Data["now_page"] = myPage.NowPage


	this.TplName="member-list.html"
}
/*
	用户列表 ui框架
 */
func(this *MemberController)List2(){
	this.TplName="member-list2.html"
}

/*
	用户列表 ui框架
 */
func(this *MemberController)List3(){
	this.TplName="member-list3.html"
}

/*
	跳添加页面
 */
func(this *MemberController)Add(){
	this.TplName="member-add.html"
}

func(this *MemberController)AddAdmin{
	name := this.GetString("name")
	if name == "" {
		//this.ReturnJson(map[string]string)
	}
}
/*
	用户列表 的 table 数据
 */
func(this *MemberController)ListData(){
	var admin []*model.Admin
	o := orm.NewOrm().QueryTable("admin")
	if startTime,endTime:=this.GetString("start"),this.GetString("end");startTime != "" && endTime != "" {
		startTime = startTime+" 00:00:00"
		endTime = endTime+" 23:59:59"
		//时间参数
		o = o.Filter("created_at__between",startTime,endTime)
	}

	if userName := this.GetString("username");userName != "" {
		//用户姓名
		o =o.Filter("name",userName)
	}
	o =o.Filter("deleted_at__isnull",true)
	o,myPage,err3:=this.GetPage(o)
	if err3 != nil {
		this.ReturnJson(map[string]string{"message":"分页错误"},400)
	}

	_,err :=o.All(&admin)
	if err != nil {
		this.ReturnJson(map[string]string{"message":"查询错误"},400)
	}

	var adminData []*model.AdminData


	for _,v := range admin {
		sexName := ""
		switch v.Sex {
			case 0:sexName = "未知"
			case 1:sexName = "男"
			case 2:sexName = "女"
			default:sexName = ""
		}
		adminData = append(adminData,&model.AdminData{v,sexName})
	}
	this.ReturnJson(map[string]interface{}{"code":0,"data":adminData,"count":myPage.Count},200)
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

	if err != nil && num > 0 {
		this.ReturnJson(map[string]string{"message":"删除失败"},400)
	}
	this.ReturnJson(map[string]string{"message":"删除成功"},200)


}
