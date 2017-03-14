package controllers

import (
)

type LogoutController struct {
	BaseController
}

func (this *LogoutController) Get() {
    this.Ctx.SetCookie("uname", "", -1, "/")
    this.Ctx.SetCookie("token", "", -1, "/")
    this.SetSession("uname","")
    this.Redirect("/", 302)
    return
}
