package models

import (
    "os"
    "path"
    "strconv"
    "strings"
    "time"
    _"github.com/Unknwon/com"
     _ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/orm"                                                                                                               
)

// 文章
type Topic struct {
    Id              int64
    Uid             int64
    Title           string
    Category        string
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


func AddTopic(title, category, lable, content, attachment string) error {
    // 处理标签
    lable = "$" + strings.Join(strings.Split(lable, " "), "#$") + "#"

    o := orm.NewOrm()

    topic := &Topic{
        Title:      title,
        Category:   category,
        Lables:     lable,
        Content:    content,
        ReplyTime:  time.Now(),
        Attachment: attachment,
        Created:    time.Now(),
        Updated:    time.Now(),
    }
    _, err := o.Insert(topic)
    if err != nil {
        return err
    }

    // 更新分类统计
    cate := new(Category)
    qs := o.QueryTable("category")
    err = qs.Filter("title", category).One(cate)
    if err == nil {
        // 如果不存在我们就直接忽略，只当分类存在时进行更新
        cate.TopicCount++
        _, err = o.Update(cate)
    }

    return err
} 

func GetTopic(tid string) (*Topic, error) {
    tidNum, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        return nil, err
    }

    o := orm.NewOrm()

    topic := new(Topic)

    qs := o.QueryTable("topic")
    err = qs.Filter("id", tidNum).One(topic)
    if err != nil {
        return nil, err
    }

    topic.Views++
    _, err = o.Update(topic)

    topic.Lables = strings.Replace(strings.Replace(
        topic.Lables, "#", " ", -1), "$", "", -1)
    return topic, nil
}

func ModifyTopic(tid, title, category, lable, content, attachment string) error {
    tidNum, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        return err
    }

    lable = "$" + strings.Join(strings.Split(lable, " "), "#$") + "#"

    var oldCate, oldAttach string
    o := orm.NewOrm()
    topic := &Topic{Id: tidNum}
    if o.Read(topic) == nil {
        oldCate = topic.Category
        oldAttach = topic.Attachment
        topic.Title = title
        topic.Category = category
        topic.Lables = lable
        topic.Content = content
        topic.Attachment = attachment
        topic.Updated = time.Now()
        _, err = o.Update(topic)
        if err != nil {
            return err
        }
    }

    // 更新分类统计
    if len(oldCate) > 0 {
        cate := new(Category)
        qs := o.QueryTable("category")
        err = qs.Filter("title", oldCate).One(cate)
        if err == nil {
            cate.TopicCount--    
            _, err = o.Update(cate)
        }
    }

    // 删除旧的附件
    if len(oldAttach) > 0 {
        os.Remove(path.Join("attachment", oldAttach))
    }

    cate := new(Category)
    qs := o.QueryTable("category")
    err = qs.Filter("title", category).One(cate)
    if err == nil {
        cate.TopicCount++
        _, err = o.Update(cate)
    }
    return nil
}

func DeleteTopic(tid string) error {
    tidNum, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        return err
    }

    o := orm.NewOrm()

    var oldCate string
    topic := &Topic{Id: tidNum}
    if o.Read(topic) == nil {
        oldCate = topic.Category
        _, err = o.Delete(topic)
        if err != nil {
            return err
        }
    }

    if len(oldCate) > 0 {
        cate := new(Category)
        qs := o.QueryTable("category")
        err = qs.Filter("title", oldCate).One(cate)
        if err == nil {
            cate.TopicCount--
            _, err = o.Update(cate)
        }
    }
    return err
}

func GetAllTopics(category, lable string, isDesc bool) (topics []*Topic, err error) {
    o := orm.NewOrm()

    topics = make([]*Topic, 0)

    qs := o.QueryTable("topic")
    if isDesc {
        if len(category) > 0 {
            qs = qs.Filter("category", category)
        }
        if len(lable) > 0 {
            qs = qs.Filter("lables__contains", "$"+lable+"#")
        }
        _, err = qs.OrderBy("-created").All(&topics)

    } else {
        _, err = qs.All(&topics)
    }
    return topics, err
}

