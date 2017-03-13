package controllers

import (
     "os"
	"strings"
    _"encoding/json"
    "strconv"
    "time"
    "fmt"
    "log"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
	"beeadmin/models"
    "encoding/base64"
    "github.com/satori/go.uuid"
    "path/filepath"
    "path"
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
	summary := this.Input().Get("summary")

	// 获取附件
	f, fh, err := this.GetFile("attachment")
    defer f.Close()
	if err != nil {
		beego.Error(err)
	}

    u1 := uuid.NewV4()
    fmt.Printf("UUIDv4: %s\n", u1)
    fmt.Printf("UUIDv4: %s\n", u1.Variant())
    fmt.Printf("UUIDv4: %s\n", u1.Version())
    /*switch v2 := u1.(type) {
        case string:
            fmt.Println(u1, "is string", v2)
        case int:
            fmt.Println(u1, "is int", v2)
        case bool:
            fmt.Println(u1, "is bool", v2)
        default:
            fmt.Println(u1, "is another type not handle yet")
    }*/
	var attachment string
	if fh != nil {
		// 保存附件
		attachment = fh.Filename
		beego.Info(attachment)
		err = this.SaveToFile("attachment", "static/upload/" + attachment)
		if err != nil {
			beego.Error(err)
		}
	}
   
   
    ff, _ := os.Open("static/upload/"+attachment)
    dir := filepath.Dir("static/upload/")
    filenameWithSuffix := path.Base(attachment)
    fmt.Println("filenameWithSuffix =", filenameWithSuffix)
    fileSuffix := path.Ext(filenameWithSuffix)
    fmt.Println("fileSuffix =", fileSuffix)
    suffix := u1.String()+fileSuffix 
    os.Rename("static/upload/"+attachment, filepath.Join(dir, suffix))
    defer ff.Close()
    sourcebuffer := make([]byte, 500000)
    n, _ := ff.Read(sourcebuffer)
    //base64压缩
    sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
   
    //fmt.Println(sourcestring)

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
    user := models.User{Name: "admin"}
    if created, id, err := o.ReadOrCreate(&user, "Name"); err == nil {
       if created {
           fmt.Println("New Insert an object. Id:", id)
       }else {
        fmt.Println("Get an object. Id:", id)
       }
    }

    var album  models.Album
    album.Owner = &user
    ids, erro := o.Insert(&album)
    if erro == nil {
        fmt.Println(ids)
    }else{
        fmt.Println(erro)
    }
    var image models.Image
    image.Thumbnail = sourcestring
    image.Created = time.Now()
    image.Owner = &user 
    image.Album = &album
    image.Uuid = u1.String()
    id, errors := o.Insert(&image)
    if errors == nil {
        fmt.Println(id)
    }else{
        fmt.Println(errors)
    }
    article := &models.Article{
        Title:      title,
        Category:   cate,
        Lables:     lable,
        Thumbnail:  &image,
        Summary:    summary,
        Content:    content,
        Created:    time.Now(),
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
    pid := article.Thumbnail.Id
    log.Println("pid - -",pid)
    image :=&models.Image{Id:pid}
    im := o.QueryTable("image") 
    err = im.Filter("id", pid).One(image)
    if err !=nil{
         log.Println("---- ",err)
    }
    this.Data["Image"] = image
    this.Data["Article"] = article
    this.TplName = "editArticle.tpl"
}

func (this *TopicController) ShowPage(){
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
    pid := article.Thumbnail.Id
    log.Println("pid - -",pid)
    image :=&models.Image{Id:pid}
    im := o.QueryTable("image") 
    err = im.Filter("id", pid).One(image)
    if err !=nil{
         log.Println("---- ",err)
    }
    this.Data["Image"] = image
    this.Data["Article"] = article
    this.TplName = "showArticle.tpl"
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
