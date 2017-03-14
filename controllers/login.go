package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
    "crypto/sha1"
    "io"
    "fmt"
    "crypto/md5"  
    "strconv"  
    "time" 
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	// 判断是否为退出操作
	if this.Input().Get("exit") == "true" {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		return
	}

	this.TplName = "login.tpl"
}

func GenToken()string{
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    return token
}

func CheckName(data string) string {
    t := sha1.New()
    io.WriteString(t, data)
    return fmt.Sprintf("%x", t.Sum(nil))
}

func (this *LoginController) Post() {
	// 获取表单信息
	uname := this.Input().Get("uname")
	pwd := CheckName(this.Input().Get("pwd"))
	autoLogin := this.Input().Get("autoLogin") == "on"
    token := GenToken()
	// 验证用户名及密码
	if uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass") {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
        this.SetSession("uname",uname)
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("token", token, maxAge, "/")
	}

	this.Redirect("/index", 302)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value
	return uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass")
}
