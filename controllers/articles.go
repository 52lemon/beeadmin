package controllers

import (
	"path"
	"strings"
    _"encoding/json"
    "strconv"
    "time"
    "fmt"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
	"beeadmin/models"
)

type TopicController struct {
	BaseController
}

func (this *TopicController)Articles(){
    this.TplName = "articles.tpl"
    pg := this.Input().Get("page")
    fmt.Println("page == %s",pg)
    page,err:= strconv.Atoi(pg)
    if err != nil {
        page =1
    }
    if page < 1 {
        page = 1
    }
    offset, _ := beego.AppConfig.Int("pageoffset")
    //todo 这里要判断前台是否传递了pageNo参数
    start := (page - 1) * offset
    o := orm.NewOrm()
    topics := make([]*models.Article, 0)
    qs := o.QueryTable("article")
    qs = qs.Limit(offset,start)
    _,err = qs.All(&topics)
    if err != nil{
      fmt.Println(err)
    }
    this.Data["Articles"] =  topics
    
}

func (this *TopicController) Save() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	// 解析表单
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	lable := this.Input().Get("lable")

	// 获取附件
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	var attachment string
	if fh != nil {
		// 保存附件
		attachment = fh.Filename
		beego.Info(attachment)
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}

    lable = "$" + strings.Join(strings.Split(lable, " "), "#$") + "#"
    o := orm.NewOrm()
    categ,er:= strconv.ParseInt(category, 10, 64)
    if er!=nil{
       beego.Error(err)
    }
    cate := &models.Category{Id:categ}
    qs := o.QueryTable("category")
    err = qs.Filter("id", cate).One(cate)
    if err != nil {
        beego.Error(err)
    }
    article := &models.Article{
        Title:      title,
        Category:   cate,
        Lables:     lable,
        Content:    content,
        ReplyTime:  time.Now(),
        Attachment: attachment,
        Created:    time.Now(),
        Updated:    time.Now(),
    }
    _, err = o.Insert(article)

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/articles", 302)
}

func (this *TopicController) AddPage(){
    o := orm.NewOrm()
    cates := make([]*models.Category, 0)
    qs := o.QueryTable("category")
    _, err := qs.All(&cates)
    this.TplName = "addArticle.tpl"
    this.Data["Categories"] = cates
    if err != nil {
        beego.Error(err)
    }
}

func (this *TopicController) EditPage(){
    tid :=this.Input().Get("id")
    tidNum, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
         beego.Error(err)
    }
    o := orm.NewOrm()
    cates := make([]*models.Category, 0)
    qs := o.QueryTable("category")
    _, err = qs.All(&cates)
    if err != nil {
        beego.Error(err)
    }
    this.Data["Categories"] = cates
    article := &models.Article{Id: tidNum}
    qss := o.QueryTable("article")
    err = qss.Filter("id", tidNum).One(article)
    if err !=nil{
         fmt.Println("---- ",err)
    }
    this.Data["Article"] = article
    this.TplName = "editArticle.tpl"

}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	tid :=this.Input().Get("tid")
    tidNum, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
         beego.Error(err)
    }

    o := orm.NewOrm()

    topic := &models.Article{Id: tidNum}
    if o.Read(topic) == nil {
        _, err = o.Delete(topic)
        if err != nil {
            beego.Error(err)
        }
    }

	this.Redirect("/topic", 302)
}

func (this *TopicController) Modify() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.TplName = "topic_modify.html"

	tid := this.Input().Get("tid")
    tidNum, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        beego.Error(err)
    }

    o := orm.NewOrm()

    topic := new(models.Article)
    qs := o.QueryTable("article")
    err = qs.Filter("id", tidNum).One(topic)
    if err != nil {
        beego.Error(err)
    }

    topic.Lables = strings.Replace(strings.Replace(
        topic.Lables, "#", " ", -1), "$", "", -1)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
	}
	this.Data["Article"] = topic
	this.Data["Tid"] = tid
	this.Data["IsLogin"] = true
}

func (this *TopicController) View() {
	this.TplName= "topic_view.html"

	reqUrl := this.Ctx.Request.RequestURI
	i := strings.LastIndex(reqUrl, "/")
	tid := reqUrl[i+1:]
    tidNum, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        beego.Error(err)
    }
    o := orm.NewOrm()
    topic := new(models.Article)
    qs := o.QueryTable("article")
    err = qs.Filter("id", tidNum).One(topic)
    if err != nil {
        beego.Error(err)
    }

    topic.Lables = strings.Replace(strings.Replace(
        topic.Lables, "#", " ", -1), "$", "", -1)
	if err != nil {
		beego.Error(err)
		this.Redirect("/articles", 302)
	}
	this.Data["Topic"] = topic
	this.Data["Lables"] = strings.Split(topic.Lables, " ")
}
