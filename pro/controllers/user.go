package controllers

import "github.com/astaxie/beego"


type UserController struct {
	beego.Controller
}

type UserInfo struct {
	Name string
	Age int
	Address string
}


func(this *UserController) Get()  {
	var userinfo UserInfo
	//name  := this.Ctx.Input.Param("name")
	name  := this.Input().Get("name")
	userinfo.Name = name
	this.Data["json"] = userinfo
	this.ServeJSON()

}