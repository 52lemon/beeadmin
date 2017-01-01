package models

import (
    "strconv"
    "time"
    _"github.com/Unknwon/com"
     _ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/orm"
)

// 分类
type Category struct {
    Id              int64
    Title           string
    Created         time.Time `orm:"index"`
    Views           int64     `orm:"index"`
    TopicTime       time.Time `orm:"index"`
    TopicCount      int64
    TopicLastUserId int64
}

func RegisterDB() {
    // 注册模型
    orm.RegisterModel(new(Category), new(Topic), new(Comment))
    // 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
    orm.RegisterDriver("mysql", orm.DRMySQL)
    // 注册默认数据库
    orm.RegisterDataBase("default","mysql","root:root@/lyblog?charset=utf8")
}

func AddCategory(name string) error {
    o := orm.NewOrm()

    cate := &Category{Title: name,Created:time.Now(),TopicTime:time.Now()}

    // 查询数据
   // qs := o.QueryTable("category")
   // err := qs.Filter("title", name).One(cate)
   // if err == nil {
   //     return err
   // }

    // 插入数据
    _, err := o.Insert(cate)
    if err != nil {
        return err
    }

    return nil
}

func DeleteCategory(id string) error {
    cid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        return err
    }

    o := orm.NewOrm()

    cate := &Category{Id: cid}
    _, err = o.Delete(cate)
    return err
}

func GetCategory(id string)(*Category,error) {
    cid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        return nil, err
    }

    o := orm.NewOrm()

    cate := &Category{Id: cid}
    qs := o.QueryTable("category")
    err = qs.Filter("id", cid).One(cate)
    if err != nil {
        return nil, err
    }
    return cate,nil 
}

func GetAllCategories() ([]*Category, error) {
    o := orm.NewOrm()

    cates := make([]*Category, 0)

    qs := o.QueryTable("category")
    _, err := qs.All(&cates)
    return cates, err
}


