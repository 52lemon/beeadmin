package routers

import (
	"beeadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginController{})
    beego.Router("/index", &controllers.HomeController{})
    beego.Router("/category", &controllers.CategoryController{})
    beego.Router("/category/add", &controllers.CategoryController{},"get:AddPage")
    beego.Router("/category/save", &controllers.CategoryController{},"post:CateAdd")
    beego.Router("/category/edit", &controllers.CategoryController{},"get:EditPage")
    beego.Router("/category/delete", &controllers.CategoryController{},"get:CateDel")
    beego.Router("/topic", &controllers.TopicController{})
    beego.AutoRouter(&controllers.TopicController{})
    beego.Router("/reply", &controllers.ReplyController{})
    beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
    beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
    beego.Router("/attachment/:all", &controllers.AttachController{})
}
