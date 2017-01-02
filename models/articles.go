package models

import (
    "time"
    _"github.com/Unknwon/com"
    _ "github.com/go-sql-driver/mysql"
    _"github.com/astaxie/beego/orm"                                                                                                               
)

// 文章
type Article struct {
    Id              int64
    Uid             int64
    Title           string
    Category        *Category   `orm:"rel(fk)"`
    Lables          string
    Content         string `orm:"size(5000)"`
    Attachment      string
    Created         time.Time `orm:"index"`
    Updated         time.Time `orm:"index"`
    Views           int64     `orm:"index"`
    Author          string
    ReplyTime       time.Time `orm:"index"`
    ReplyCount      int64
    ReplyLastUserId int64
}



