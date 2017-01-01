package controllers

import (
	_"github.com/astaxie/beego"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) Get() {
	this.TplName = ""
}
