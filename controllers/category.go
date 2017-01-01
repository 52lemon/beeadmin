package controllers

import (
	"beeadmin/models"
	"github.com/astaxie/beego"
    "log"
)

type CategoryController struct {
	BaseController
}

//获取并显示分类列表
func (this *CategoryController) Get() {
	// 检查是否有操作
    cates,err :=models.GetAllCategories()

	this.TplName = "category.tpl"
	this.Data["Categories"] = cates
	if err != nil {
		beego.Error(err)
	}
}

func (this *CategoryController)AddPage(){
     this.TplName ="addCategory.tpl"
}

func (this *CategoryController)EditPage(){
     id := this.Input().Get("id")
     cate,err := models.GetCategory(id)
     if err !=nil{
         beego.Error(err)
     }
     this.Data["Category"] = cate
     this.TplName ="editCategory.tpl"
}

func (this *CategoryController)CateAdd(){
        name := this.Input().Get("name")

        err := models.AddCategory(name)
        if err != nil {
            beego.Error(err)                                                                                                                     
        }

        this.Redirect("/category", 302)
        return
}

func (this *CategoryController)CateDel(){
     id := this.Input().Get("id") 
     log.Println(id)
     err:= models.DeleteCategory(id)
     if err !=nil{
        this.Data["json"] = models.NewErrorInfo("Error")
        this.ServeJSON()
        return
     }
     this.Data["json"] = models.NewNormalInfo("Succes")
     this.ServeJSON()
}
