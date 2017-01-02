package controllers

import (
	"beeadmin/models"
    "time"
    "strconv"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
)

type ReplyController struct {
	BaseController
}

func (this *ReplyController) Add() {
	tid := this.Input().Get("tid")
    tidNum, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        beego.Error(err)
    }

    reply := &models.Comment{
        Tid:     tidNum,
        Name:    this.Input().Get("nickname"),
        Content: this.Input().Get("content"),
        Created: time.Now(),
    }
    o := orm.NewOrm()
    _, err = o.Insert(reply)
    if err != nil {
		beego.Error(err)
    }

    topic := &models.Article{Id: tidNum}
    if o.Read(topic) == nil {
        topic.ReplyTime = time.Now()
        topic.ReplyCount++
        _, err = o.Update(topic)
    }
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ReplyController) Delete() {
	if !checkAccount(this.Ctx) {
		return
	}
	tid := this.Input().Get("tid")
    ridNum, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        beego.Error(err)
    }
    o := orm.NewOrm()
    var tidNum int64
    reply := &models.Comment{Id: ridNum}
    if o.Read(reply) == nil {
        tidNum = reply.Tid
        _, err = o.Delete(reply)
        if err != nil {
          beego.Error(err)
        }
    }
    replies := make([]*models.Comment, 0)
    qs := o.QueryTable("comment")
    _, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
    if err != nil {
        beego.Error(err)
    }

    topic := &models.Article{Id: tidNum}
    if o.Read(topic) == nil {
        topic.ReplyTime = replies[0].Created
        topic.ReplyCount = int64(len(replies))
        _, err = o.Update(topic)
    }
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}
