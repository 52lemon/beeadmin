package controllers

import (
	"beeadmin/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "index.tpl"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(
		this.Input().Get("cate"), this.Input().Get("lable"), true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
}
