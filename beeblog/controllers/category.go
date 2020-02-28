package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c * CategoryController) Get() {
	if!checkAccount(c.Ctx){
		c.Redirect("/login",302)
	}

	c.Data["IsLogin"]=checkAccount(c.Ctx)
	op:=c.Input().Get("op")
	switch op{
	case "add":
		name:=c.Input().Get("name")
		if len(name)==0{
			break
		}
		err :=models.AddCategory(name)
		if err!=nil{
			beego.Error(err)
		}
		c.Redirect("/category",301)
		return
	case "del":
		id:=c.Input().Get("id")
		fmt.Println(id)
		if len(id)==0{
			break
		}
		err:=models.DelCategory(id)
		if err!=nil{
			beego.Error(err)
		}
		c.Redirect("/category",301)
		return
	}
	c.TplName = "category.html"
    c.Data["IsCategory"]=true
    var err error
    c.Data["Categories"],err=models.GetAllCategories()
    if err!=nil{
    	beego.Error(err)
	}
}
