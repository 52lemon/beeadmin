package main

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _"beeadmin/routers"
	_"beeadmin/controllers"
	"beeadmin/models"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func makeUrl(source string) string { 
    return  "/static/upload/"+source
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	// 附件处理
	os.Mkdir("attachment", os.ModePerm)
   
    beego.AddFuncMap("makeUrl",makeUrl)
	// 启动 beego
	beego.Run()
}
