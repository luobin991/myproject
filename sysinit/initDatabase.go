package sysinit

import (
	_ "myproject/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase(){
	//读取配置文件，设置数据库参数
	dbType 	:= beego.AppConfig.String("db_type")			//数据库类别
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")//连接名称
	dbName 	:= beego.AppConfig.String(dbType + "::db_name")	//数据库名称
	dbUser 	:= beego.AppConfig.String(dbType + "::db_user")	//数据库连接用户名
	dbPwd 	:= beego.AppConfig.String(dbType + "::db_pwd")	//数据库连接用户名
	dbHost 	:= beego.AppConfig.String(dbType + "::db_host")	//数据库IP（域名）
	dbPort 	:= beego.AppConfig.String(dbType + "::db_port")	//数据库端口

	switch dbType {
	case "sqlite3":
		orm.RegisterDataBase(dbAlias,dbType,dbName)
	case "mysql":
		dbCharset := beego.AppConfig.String(dbType + "::db_charset")
		orm.RegisterDataBase(dbAlias,dbType,dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" +dbPort + ")/" + dbName + "?charset=" + dbCharset,30)
	}
	isDev := (beego.AppConfig.String("runmode") == "dev") 	//如果是开发模式，则显示命令信息
	orm.RunSyncdb("default",false,isDev) 	//自动建表
	if isDev {
		orm.Debug = isDev
	}
}