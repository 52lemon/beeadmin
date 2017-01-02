package models

import (
    "time"
    _"github.com/Unknwon/com"
     _ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/orm"
)

// 分类
type Category struct {
    Id              int64
    Title           string
    Created         time.Time   `orm:"index"`
    Thumbnail       string
    Description     string
    Parent          *Category   `orm:"rel(fk)"`
    Article         []*Article  `orm:"reverse(many)"`
}

func RegisterDB() {
    // 注册模型
    orm.RegisterModel(new(Category), new(Article), new(Comment))
    // 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
    orm.RegisterDriver("mysql", orm.DRMySQL)
    // 注册默认数据库
    orm.RegisterDataBase("default","mysql","root:root@/beeadmin?charset=utf8")
}
