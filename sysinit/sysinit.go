//sysinit.go
package sysinit
import(
	"myproject/utils"
	"github.com/astaxie/beego"
	)

func init() {
	beego.BConfig.WebConfig.Session.SessionOn = true //启用Session
	utils.InitLogs()		//初始化日志
	utils.InitCache()		//初始化缓存
	InitDatabase() 			//初始化数据库
}