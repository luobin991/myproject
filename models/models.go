package models
import (
		"os"
		"path"
		"time"
		"github.com/Unknwon/com"
		"github.com/astaxie/beego/orm"
		_ "github.com/go-sql-driver/mysql" 
	)

const(
	_DB_NAME = "root:ly123@tcp(192.168.0.150:3306)/test?charset=utf8" //ly123
	_mysql_DRIVER = "mysql"
)
type Category struct{
	Id 				int64
	Title 			string
	Created 		time.Time `orm:"index"`
	Views 			int64 `orm:"index"`
	TopicTime 		time.Time `orm:"index"`
	TopicCount 		int64
	TopicLastUserId int64
	}

type Topic struct{
	Id 				int64
	Uid 			int64
	Title 			string
	Content 		string `orm:"size(5000)"`
	Attachment 		string
	Updated 		time.Time `orm:"index"`
	Created 		time.Time `orm:"index"`
	Views 			int64 `orm:"index"`
	Autthor 		string
	ReplyTime 		time.Time `orm:"index"`
	ReplyCount 		int64
	ReplyLastUserId int64
}


func RegisterDB(){
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category),new(Topic))
	orm.RegisterDriver(_mysql_DRIVER,orm.DRMySQL)
	orm.RegisterDataBase("default",_mysql_DRIVER,_DB_NAME,10)

}