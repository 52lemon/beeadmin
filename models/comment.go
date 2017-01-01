package models

import (
    "strconv"
    "time"
    _"github.com/Unknwon/com"
     _ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/orm"
)

// 评论
type Comment struct {
    Id      int64
    Tid     int64
    Name    string
    Content string    `orm:"size(1000)"`
    Created time.Time `orm:"index"`
}

func AddReply(tid, nickname, content string) error {
    tidNum, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        return err
    }

    reply := &Comment{
        Tid:     tidNum,
        Name:    nickname,
        Content: content,
        Created: time.Now(),
    }
    o := orm.NewOrm()
    _, err = o.Insert(reply)
    if err != nil {
        return err
    }

    topic := &Topic{Id: tidNum}
    if o.Read(topic) == nil {
        topic.ReplyTime = time.Now()
        topic.ReplyCount++
        _, err = o.Update(topic)
    }
    return err
}

func GetAllReplies(tid string) (replies []*Comment, err error) {
    tidNum, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        return nil, err
    }

    replies = make([]*Comment, 0)

    o := orm.NewOrm()
    qs := o.QueryTable("comment")
    _, err = qs.Filter("tid", tidNum).All(&replies)
    return replies, err
}

func DeleteReply(rid string) error {
    ridNum, err := strconv.ParseInt(rid, 10, 64)
    if err != nil {
        return err
    }

    o := orm.NewOrm()

    var tidNum int64
    reply := &Comment{Id: ridNum}
    if o.Read(reply) == nil {
        tidNum = reply.Tid
        _, err = o.Delete(reply)
        if err != nil {
            return err
        }
    }

    replies := make([]*Comment, 0)
    qs := o.QueryTable("comment")
    _, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
    if err != nil {
        return err
    }

    topic := &Topic{Id: tidNum}
    if o.Read(topic) == nil {
        topic.ReplyTime = replies[0].Created
        topic.ReplyCount = int64(len(replies))
        _, err = o.Update(topic)
    }
    return err
}

