package controllers

import (
	// "strconv"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	isLogin := checkAccount(this.Ctx)
	
	beego.Info("index get IsLogin value：",isLogin)

	this.Data["IsLogin"]  = isLogin
	this.Data["IsLogins"] = [1]bool{isLogin}

	this.TplName = "index.html"
	
}



















/*
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me."
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	c.Ctx.WriteString("AppName:"+beego.AppConfig.String("appname")+
		"\nhttpPort:"+beego.AppConfig.String("httpport")+
		"\nrunMode:"+beego.AppConfig.String("runmode"))

	hp := strconv.Itoa(beego.BConfig.Listen.HTTPPort)
	c.Ctx.WriteString("\n\nAppName:" + beego.BConfig.AppName +
		"\nhttpPort:" + hp +
		"\nrunMode:" + beego.BConfig.RunMode)

	beego.Trace("trace test1")
	beego.Info("info test1")

	beego.SetLevel(beego.LevelInformational)
	
	beego.Trace("trace test2")
	beego.Info("info test2")

	c.Data["TrueCond"] = true
	c.Data["FalseCond"] = false

	type u struct {
		Name string
		Age int
		Sex string
	}
	user := &u{
		Name:"joc",
		Age:20,
		Sex:"Male",
	}
	c.Data["User"] = user

	nums := []int{1,2,3,4,5,6,7,8,9,0}
	c.Data["Nums"] = nums

	c.Data["TplVar"]="hey guys"
	c.Data["Html"] = "<div><stong>Say hello Go html</stong></div>"

	
}




  <!-- <div class="description">
    <div>
      {{if .TrueCond}}
        true condition.
      {{end}}
    </div>
    <div>
      {{if .FalseCond}}
        false condition.
      {{end}}
    </div>

    <div>
      {{with .User}}
       name:{{.Name}}  <strong>age</strong>:{{.Age}} sex:{{.Sex}}
      {{end}}
    <div>
      {{range .Nums}}
        {{.}}
      {{end}}
    </div>

    <div>
      {{$tplVar := .TplVar}}
      {{$tplVar}}
    </div>

    {{.Html
    {{str2html .Html
    {{.Html | htmlquote

    </div>
      {{template "test"}}
  </div> -->
*/