package routers

import (
	"beeadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginController{})
    beego.Router("/logout", &controllers.LogoutController{})
    beego.Router("/index", &controllers.HomeController{})
    beego.Router("/category", &controllers.CategoryController{})
    beego.Router("/category/add", &controllers.CategoryController{},"get:AddPage")
    beego.Router("/category/save", &controllers.CategoryController{},"post:CateAdd")
    beego.Router("/category/edit", &controllers.CategoryController{},"get:EditPage")
    beego.Router("/category/modify", &controllers.CategoryController{},"post:Modify")
    beego.Router("/category/delete", &controllers.CategoryController{},"get:CateDel")
    beego.Router("/ctree/get", &controllers.CategoryController{},"get:TreeGet")
    beego.Router("/ctree/delete", &controllers.CategoryController{},"post:TreeDel")
    beego.Router("/ctree/save", &controllers.CategoryController{},"post:TreeSave")
    beego.Router("/ctree/modify", &controllers.CategoryController{},"post:TreeModify")
    beego.Router("/articles", &controllers.TopicController{},"get:Articles")
    beego.Router("/article/add", &controllers.TopicController{},"get:AddPage")
    beego.Router("/article/save", &controllers.TopicController{},"post:Save")
    beego.Router("/article/edit", &controllers.TopicController{},"get:EditPage")
    beego.Router("/article/show", &controllers.TopicController{},"get:ShowPage")
    beego.Router("/article/modify", &controllers.TopicController{},"post:Modify")
    beego.Router("/reply", &controllers.ReplyController{})
    beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
    beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
    beego.Router("/attachment/:all", &controllers.AttachController{})
}
