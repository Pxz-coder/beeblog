package routers

import (
	"beeblog/controllers"
	"beeblog/models"
	"github.com/astaxie/beego"
	"os"
)

func init() {
	models.RegisterDB()
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/category",&controllers.CategoryController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/topic",&controllers.TopicController{})
	beego.Router("/reply",&controllers.ReplyController{})
	beego.Router("/reply/add",&controllers.ReplyController{},"post:Add")
	beego.Router("/reply/delete",&controllers.ReplyController{},"get:Delete")
	//作为静态文件
	os.Mkdir("attachment",os.ModePerm)
	beego.SetStaticPath("/attachment","attachment")

}
