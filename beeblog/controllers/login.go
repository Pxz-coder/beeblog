package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct{
	beego.Controller
}

func (c *LoginController) Get() {

	isExit:=c.Input().Get("exit")=="true"
	if isExit{
		c.Ctx.SetCookie("uname","",-1,"/")
		c.Ctx.SetCookie("pwd","",-1,"/")
		c.Redirect("/",301)
		return
	}
	c.TplName = "login.html"

}

func (c *LoginController) Post() {

	uname:=c.Input().Get("uname")
	pwd:=c.Input().Get("pwd")
	autoLogin:=c.Input().Get("autoLogin")=="on"
	if beego.AppConfig.String("uname")==uname&&
		beego.AppConfig.String("pwd")==pwd{
		maxAge:=0
		if autoLogin{
			maxAge=1<<31-1
		}
		c.Ctx.SetCookie("uname",uname,maxAge,"/")
		c.Ctx.SetCookie("pwd",pwd,maxAge,"/")
		c.Redirect("/",301)
	}
	return
}

func checkAccount(ctx *context.Context) bool{
	ck,err:=ctx.Request.Cookie("uname")
	if err!=nil{
		return false
	}
	uname:=ck.Value
	ck,err=ctx.Request.Cookie("pwd")
	if err!=nil{
		return false
	}
	pwd:=ck.Value

	return beego.AppConfig.String("uname")==uname&&
		beego.AppConfig.String("pwd")==pwd
}
