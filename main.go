package main

import (
	_ "myproject/routers"
	_ "myproject/sysinit"

	"github.com/astaxie/beego"
)

func main() {

	beego.Run()

	//https://sourcegraph.com/github.com/lhtzbj12/sdrms@master/-/blob/utils/cache.go#L57
}

