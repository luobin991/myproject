package controllers

import (
	// "net/url"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	)

type LoginController struct {
	beego.Controller
}


func (this *LoginController) Get() {
	input := this.Input().Get("exit")
	isExit := input == "0"

	beego.Info("login isExit:",input) 
	
	if isExit {
		this.Ctx.SetCookie("uname","",-1,"/")
		this.Ctx.SetCookie("pwd","",-1,"/")
		this.Redirect("/",301)
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post(){
	// this.Ctx.WriteString(fmt.Sprint(this.Input()))

	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"


		beego.Info("uname：",uname) 
		// beego.Info("pwd：",pwd)
		beego.Info("autoLogin：",autoLogin)

		// beego.Info("AppConfig uname：",beego.AppConfig.String("uname")) 
		// beego.Info("AppConfig pwd：",beego.AppConfig.String("pwd") )

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
			maxAge := 0
			beego.Info("Login Ok")
			if autoLogin {
				maxAge = 1<<31 - 1
			}
			this.Ctx.SetCookie("uname",uname,maxAge,"/")
			this.Ctx.SetCookie("pwd",pwd,maxAge,"/")
		}

	this.Redirect("/",301)
	return
}


func checkAccount(ctx *context.Context) bool{

	uname := ctx.GetCookie("uname")
	  
	 pwd := ctx.GetCookie("pwd")
	 
	checkAccountIsLogin := beego.AppConfig.String("uname") ==uname &&
		beego.AppConfig.String("pwd") == pwd

	
	beego.Info("checkAccountIsLogin：",checkAccountIsLogin)
	return checkAccountIsLogin	

}