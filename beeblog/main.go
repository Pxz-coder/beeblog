package main

import (
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


func main() {

	orm.RunSyncdb("default",true,false)
	beego.Run()
}

