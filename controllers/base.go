package controllers

import (
    "github.com/astaxie/beego"
)

type BaseController struct {
    beego.Controller
}

// Predefined const error strings.
const (
    ErrInputData    = "数据输入错误"
    ErrDatabase     = "数据库操作错误"
    ErrDupUser      = "用户信息已存在"
    ErrNoUser       = "用户信息不存在"
    ErrPass         = "密码不正确"
    ErrNoUserPass   = "用户信息不存在或密码不正确"
    ErrNoUserChange = "用户信息不存在或数据未改变"
    ErrInvalidUser  = "用户信息不正确"
    ErrOpenFile     = "打开文件出错"
    ErrWriteFile    = "写文件出错"
    ErrSystem       = "操作系统错误"
)

// ControllerError is controller error info structer.
type ControllerError struct {
    Status   int    `json:"status"`
    Code     int    `json:"code"`
    Message  string `json:"message"`
    DevInfo  string `json:"dev_info"`
    MoreInfo string `json:"more_info"`
}

// Predefined controller error values.
var (
    err404          = &ControllerError{404, 404, "page not found", "page not found", ""}
    errInputData    = &ControllerError{400, 10001, "数据输入错误", "客户端参数错误", ""}
    errDatabase     = &ControllerError{500, 10002, "服务器错误", "数据库操作错误", ""}
    errDupUser      = &ControllerError{400, 10003, "用户信息已存在", "数据库记录重复", ""}
    errNoUser       = &ControllerError{400, 10004, "用户信息不存在", "数据库记录不存在", ""}
    errPass         = &ControllerError{400, 10005, "用户信息不存在或密码不正确", "密码不正确", ""}
    errNoUserPass   = &ControllerError{400, 10006, "用户信息不存在或密码不正确", "数据库记录不存在或密码不正确", ""}
    errNoUserChange = &ControllerError{400, 10007, "用户信息不存在或数据未改变", "数据库记录不存在或数据未改变", ""}
    errInvalidUser  = &ControllerError{400, 10008, "用户信息不正确", "Session信息不正确", ""}
    errOpenFile     = &ControllerError{500, 10009, "服务器错误", "打开文件出错", ""}
    errWriteFile    = &ControllerError{500, 10010, "服务器错误", "写文件出错", ""}
    errSystem       = &ControllerError{500, 10011, "服务器错误", "操作系统错误", ""}
    errExpired      = &ControllerError{400, 10012, "登录已过期", "验证token过期", ""}
    errPermission   = &ControllerError{400, 10013, "没有权限", "没有操作权限", ""}
)

// RetError return error information in JSON.
func (base *BaseController) RetError(e *ControllerError) {
    if mode := beego.AppConfig.String("runmode"); mode == "prod" {
        e.DevInfo = ""
    }

    base.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
    base.Ctx.ResponseWriter.WriteHeader(e.Status)
    base.Data["json"] = e
    base.ServeJSON()

    base.StopRun()
}

