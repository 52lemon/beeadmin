package controllers

import (
    _"beeadmin/models"

    _"github.com/astaxie/beego"
)

type HomeController struct {
    BaseController
}

func (this *HomeController) Get() {
    this.Layout = "base.html"
    this.TplName = "index.html"
}
