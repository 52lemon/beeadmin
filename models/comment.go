package models

import (
    "time"
    _"github.com/Unknwon/com"
    _ "github.com/go-sql-driver/mysql"
    _"github.com/astaxie/beego/orm"
)

// 评论
type Comment struct {
    Id      int64
    Tid     int64
    Name    string
    Content string    `orm:"size(1000)"`
    Created time.Time `orm:"index"`
}


