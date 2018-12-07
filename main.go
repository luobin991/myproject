package main

import (
	_ "myproject/routers"
	_ "myproject/controllers"
	"myproject/models"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)


func init() {
	models.RegisterDB()
}


func main() {
	orm.Debug = true
	// o := orm.NewOrm()
	// o.Using("qianbao_test")
	orm.RunSyncdb("default",false,true)
	// beego.Router("/",&controllers.MainController{})

	beego.Run()

	//https://sourcegraph.com/github.com/lhtzbj12/sdrms@master/-/blob/utils/cache.go#L57
}

