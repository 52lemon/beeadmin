package controllers

import (
	"io"
	"net/url"
	"os"

	_"github.com/astaxie/beego"
)

type AttachController struct {
	BaseController
}

func (this *AttachController) Get() {
	filePath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()

	_, err = io.Copy(this.Ctx.ResponseWriter, f)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
}
