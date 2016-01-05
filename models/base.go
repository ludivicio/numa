package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// dbhost := beego.AppConfig.String("dbhost")
	// dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")

	// if dbport == "" {
	// 	dbport = "3306"
	// }

	//注册mysql Driver
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	//构造conn连接
	conn := dbuser + ":" + dbpassword + "@/" + dbname + "?charset=utf8"
	//数据库别名
	alice := "default"
	//不强制建数据库
	force := false
	//打印建表过程
	verbose := true

	orm.RegisterDataBase(alice, "mysql", conn)
	orm.RegisterModel(new(Admin), new(Article), new(ArticleCate), new(AutoUser), new(Baoliao), new(Item), new(ItemCate), new(ItemComment), new(ItemTag), new(Mall), new(Tag))
	//建表
	err := orm.RunSyncdb(alice, force, verbose)
	if err != nil {
		beego.Error(err)
	}

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

}

// TableName 返回带前缀的表名
func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}
