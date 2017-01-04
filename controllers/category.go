package controllers

import (
    "strconv"
	"beeadmin/models"
	"github.com/astaxie/beego"
    "log"
    "time"
    "fmt"
    "github.com/astaxie/beego/orm"
)

type CategoryController struct {
	BaseController
}

//获取并显示分类列表
func (this *CategoryController) Get() {
	// 检查是否有操作
    o := orm.NewOrm()
    cates := make([]*models.Category, 0)
    qs := o.QueryTable("category")
    _, err := qs.All(&cates)
	if err != nil {
		beego.Error(err)
	}
	this.TplName = "category.tpl"
	this.Data["Categories"] = cates
}

func (this *CategoryController)AddPage(){
    this.TplName ="addCategory.tpl"
    o := orm.NewOrm()
    cates := make([]*models.Category, 0)
    qs := o.QueryTable("category")
    _, err := qs.All(&cates) 
    this.Data["Categories"] = cates
    if err != nil {
        fmt.Println(err)
        beego.Error(err)
    }
}

func (this *CategoryController)EditPage(){
     id := this.Input().Get("id")
     cid, err := strconv.ParseInt(id, 10, 64)
     if err != nil {
        beego.Error(err)
     }
    o := orm.NewOrm()
    cates := make([]*models.Category, 0)
    qs := o.QueryTable("category")
    _, err = qs.All(&cates) 
    this.Data["Categories"] = cates
     cate := &models.Category{Id: cid}
     qss := o.QueryTable("category")
     err = qss.Filter("id", cid).One(cate)
     if err !=nil{
         beego.Error(err)
     }
     this.Data["Category"] = cate
     this.TplName ="editCategory.tpl"
}

func (this *CategoryController)Modify(){
     id := this.Input().Get("id")
     category := this.Input().Get("category")
     title := this.Input().Get("title")
     desc := this.Input().Get("desc")
     catid, err := strconv.ParseInt(category, 10, 64)
     cat := &models.Category{Id: catid}
     if err != nil {
        fmt.Println(err)
        beego.Error(err)
     }
     o := orm.NewOrm()
     qss := o.QueryTable("category")
     err = qss.Filter("id", catid).One(cat)
     if err !=nil{
        fmt.Println(err)
         beego.Error(err)
     }
     cid, err := strconv.ParseInt(id, 10, 64)
     if err != nil {
        fmt.Println(err)
        beego.Error(err)
     }
    cates := make([]*models.Category, 0)
    qs := o.QueryTable("category")
    _, err = qs.All(&cates) 
    this.Data["Categories"] = cates
    cate := &models.Category{Id: cid}
    fmt.Println("cid == %d",cid)
    if o.Read(cate) == nil {
        cate.Title =title
        cate.Parent = cat
        cate.Description =  desc
        num, er := o.Update(cate,"Title","Parent","Description")
        if err == nil {
            fmt.Println("---- %d",num)
       }else{
          fmt.Println(er)
      }
    }  
     this.Redirect("/category", 302)
}

func (this *CategoryController)CateAdd(){
        name := this.Input().Get("name")
        des := this.Input().Get("desc")
        o := orm.NewOrm()
        id := this.Input().Get("cid")
        cid, err := strconv.ParseInt(id, 10, 64)
        cate := &models.Category{Id: cid}
        qs := o.QueryTable("category")
        err = qs.Filter("id", cid).One(cate)
        if err != nil {
            beego.Error(err)
        }
        fmt.Println(cate.Id)
        fmt.Println(cate.Title)
        category := &models.Category{Title: name,Created:time.Now(),Parent:cate,Description:des}
    // 查询数据
   // qs := o.QueryTable("category")
   // err := qs.Filter("title", name).One(cate)
   // if err == nil {
   //     return err
   // }

    // 插入数据
        _, err = o.Insert(category)
        if err != nil {
            beego.Error(err) 
        }
        this.Redirect("/category", 302)
}

func (this *CategoryController)CateDel(){
     id := this.Input().Get("id") 
     log.Println(id)
     cid, err := strconv.ParseInt(id, 10, 64)
     if err != nil {
        beego.Error(err)
     }
     o := orm.NewOrm()
     cate := &models.Category{Id: cid}
     _, err = o.Delete(cate)
     if err !=nil{
        this.Data["json"] = models.NewErrorInfo("Error")
        this.ServeJSON()
        return
     }
     this.Data["json"] = models.NewNormalInfo("Succes")
     this.ServeJSON()
}
