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
	this.TplName = "category.tpl"
	this.Data["Categories"] = cates
	if err != nil {
		beego.Error(err)
	}
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
     cate := &models.Category{Id: cid}
     qs := o.QueryTable("category")
     err = qs.Filter("id", cid).One(cate)
     if err !=nil{
         beego.Error(err)
     }
     this.Data["Category"] = cate
     this.TplName ="editCategory.tpl"
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
