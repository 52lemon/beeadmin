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
    Product         []*Product  `orm:"reverse(many)"`
}
// 图片
type Album struct {
    Id              int64
    Description     string  
    thumbnail       string
    Owner           *User          `orm:"reverse(one)"`
    Group           *Group         `orm:"reverse(one)"`
    Product         *Product       `orm:"reverse(one)"`
    Image           []*Image       `orm:"reverse(many)"`
}

// 图片
type Image struct {
    Id              int64
    Uuid            string
    Link            string
    Thumbnail       string
    Created         time.Time    `orm:"index"`
    Description     string  
    Owner           *User        `orm:"rel(fk)"`
    Album           *Album       `orm:"rel(fk)"` 
}

// 用户
type User struct {
    Id              int64
    Name            string
    Cellphone       string
    password        string
    thumbnail       *Image      `orm:"rel(one)"`
    Created         time.Time    `orm:"index"`
    Logined         time.Time    `orm:"index"`
    Email           string
    Address         string
    Avatur          string
    Gender          int
    Role            int
    Token           string
    Tokened         time.Time    `orm:"index"`
    Description     string  
    Album           *Album     `orm:"rel(one)"`                                                                                                  
}

// 文章
type Article struct {
    Id              int64
    Title           string
    Category        *Category    `orm:"rel(fk)"`
    Summary        string
    Thumbnail       *Image      `orm:"rel(one)"`                                                                                                 
    Lables          string
    Content         string       `orm:"size(5000)"`
    Created         time.Time    `orm:"index"`
    Views           int64        `orm:"index"`
    Author          string
    Origin          string
}

// 评论
type Comment struct {                                                                                                                            
    Id      int64
    Tid     int64
    Name    string
    Content string    `orm:"size(1000)"`
    Created time.Time `orm:"index"`
}

// 公司
type Group struct {                                                                                                                              
    Id              int64
    Name            string
    Link            string
    Summmary        string
    thumbnail       string
    Created         time.Time    `orm:"index"`
    Description     string  
    Address         string
    Email           string
    Postcode        string
    Contact         string
    Album           *Album       `orm:"rel(one)"`                                                                                                
}
// 产品
type Product struct {
    Id              int64
    Name            string
    Price           float32 
    Discount        float32 
    Description     string
    Summary         string 
    Category        *Category    `orm:"rel(fk)"` 
    thumbnail       string
    Album           *Album        `orm:"rel(one)"`                                                                                               
    Created         time.Time    `orm:"index"`
}

func RegisterDB() {
    // 注册模型
    orm.RegisterModel(new(Category), new(Article), new(Comment),new(Product),new(Group),new(Image),new(Album),new(User))                         
    // 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
    orm.RegisterDriver("mysql", orm.DRMySQL)
    // 注册默认数据库
    orm.RegisterDataBase("default","mysql","root:root@/beeadmin?charset=utf8")
}
