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

type Tree struct {
    Id       int    `json:"id"`
    Parent   int    `json:"parent"`
    Text     string `json:"text"`
    Haschildren bool   `json:"haschildren"`
    Hasbrother  bool   `json:"hasbrother"`
    Isend       bool  `json:"isend"` 
    Deep       int  `json:"deep"` 
    Hero       []int `json:"hero"`
}

//前序遍历分类构建分类树数据结构
func TreeList(id int,deep int)(treelist []*Tree){
    ts := make([]*Tree, 0, 0)
    o := orm.NewOrm()
    var leaf Tree
    cate := models.Category{Id: int64(id)}
    err := o.Read(&cate)
    if err == orm.ErrNoRows {
        fmt.Println("查询不到")
    } else if err == orm.ErrMissPK {
        fmt.Println("找不到主键")
    } 
    leaf.Id =  int(cate.Id)
    leaf.Text = cate.Title
    leaf.Haschildren = false
    leaf.Hasbrother = false
    leaf.Isend = false
    leaf.Deep = deep
    if cate.Id == 5{
       leaf.Parent = 0
    }else{
       leaf.Parent = int(cate.Parent.Id)
    }
    cates := make([]*models.Category, 0)
    qs := o.QueryTable("category")
    _, err = qs.Filter("parent_id",cate.Id).All(&cates)
    if err != nil {
        beego.Error(err)
    }
    cnt,_:= o.QueryTable("category").Filter("parent_id",cate.Id).Count()
    if cnt >0{
        leaf.Haschildren = true
    }
    ts = append(ts,&leaf)
    for i:=0;i<len(cates);i++{
        for _,v := range(TreeList(int(cates[i].Id),int(deep+1))){
            ts = append(ts,v) 
       }
    }
    return ts
}

func CacDeep(lst []*Tree,index int)int{
    deep := lst[index].Deep
    for i,value := range lst{
        if (value.Id == lst[index].Parent && value.Hasbrother && !value.Isend) || (value.Id == lst[index].Parent && value.Id==5){
            deep = value.Deep
        }else if value.Id==lst[index].Parent{
            deep = CacDeep(lst,i)
        }
    }
    return deep
}

func CheckIsend(lst []*Tree,index int){
    for i,v := range lst{
        if v.Id == lst[index].Parent{
            CheckIsend(lst,i)
            v.Isend  =  false
        }
    }
}

func Trees()(treelist []*Tree){
    tree := make([]*Tree,0,0)
    isends := make([]*Tree,0,0)
    sortid := make([]int,0,0)
    tree = TreeList(5,0)
    // 计算是否是end节点,是否有兄弟
    for _,value := range tree{
        for i,v := range tree{
            if value.Id == v.Parent{
                sortid= append(sortid,i)
            }
        }
        if lastindex:=len(sortid);lastindex>0{
            index := lastindex -1
            tree[sortid[index]].Isend = true
        }
        if lastindex:=len(sortid);lastindex>1{
            for _,val := range sortid{
                tree[val].Hasbrother = true
            }
        }
    }
    for _,vl := range tree{
       if vl.Isend == true{
          isends = append(isends,vl)
       } 
    }
    //计算补偿数
    for in,va := range tree{
        if va.Isend==true{
            deep := CacDeep(tree,in)
            va.Deep = va.Deep - deep
        }
    }
    for _,valu := range isends{
        if valu.Isend==true{
            for _,ve := range isends{
                if valu.Id == ve.Parent{
                    valu.Isend = false 
                }
            }
        }
    }
    for _,va := range tree{
       if va.Isend{
           hero := make([]int,va.Deep)
           va.Hero = hero
           for z :=0;z<va.Deep;z++{
               va.Hero[z] = va.Deep
           }
       }
    }
    return tree
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
    lists := make([]*Tree,0,0)
    lists = Trees() 
    fmt.Println("------",lists)
	this.TplName = "category.html"
    fmt.Println("category is",cates)
	this.Data["Categories"] = cates
	this.Data["Trees"] = lists
}

func (this *CategoryController)AddPage(){
    this.TplName ="addCategory.html"
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
     this.TplName ="editCategory.html"
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

func (this *CategoryController)TreeDel(){
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

func (this *CategoryController)TreeGet(){
     id := this.Input().Get("id") 
     log.Println(id)
     cid, err := strconv.ParseInt(id, 10, 64)
     if err != nil {
         beego.Error(err)
     }
     cate := models.Category{Id: cid}
     o := orm.NewOrm()
     err = o.Read(&cate)
     if err == orm.ErrNoRows {
        fmt.Println("查询不到")
     } else if err == orm.ErrMissPK {
        fmt.Println("找不到主键")
     } 
     this.Data["json"] = cate
     this.ServeJSON()
}

func (this *CategoryController)TreeSave(){
     name:= this.Input().Get("name") 
     p_id:= this.Input().Get("parent") 
     //log.Println(id)
     o := orm.NewOrm()
     cid, err := strconv.ParseInt(p_id, 10, 64)
     if err != nil {
         beego.Error(err)
     }
     parent := models.Category{Id: cid}
     err = o.Read(&parent)
     if err != nil {
        beego.Error(err)
     }
     category := models.Category{Title: name,Created:time.Now(),Parent:&parent,Description:name}  
     id, err := o.Insert(&category)
     if err == nil {
        fmt.Println(id)
        this.Data["json"] =  models.NewNormalInfo("Succes")
     }else{
       
       this.Data["json"] = models.NewNormalInfo("Error")
     }
     this.ServeJSON()
}

func (this *CategoryController)TreeModify(){
     name:= this.Input().Get("name") 
     c_id:= this.Input().Get("id") 
     //log.Println(id)
     cid, err := strconv.ParseInt(c_id, 10, 64)
     if err != nil {
            beego.Error(err)
     }
     o := orm.NewOrm()
     num, erro := o.QueryTable("category").Filter("id", cid).Update(orm.Params{
        "title": name,
     })
     if erro != nil {
        this.Data["json"] = models.NewNormalInfo("Error")
     }else{
        this.Data["json"] =  models.NewNormalInfo("Succes")
        fmt.Printf("Affected Num: %s, %s", num, err)
     }
     this.ServeJSON()
}

