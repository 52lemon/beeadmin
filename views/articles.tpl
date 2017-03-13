<!DOCTYPE html>
<html lang="en">
<head>
    <title>文章列表</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <link rel="icon" href="favicon.ico" type="image/x-icon" />
    <link rel="stylesheet" type="text/css" id="theme" href="../static/joli/css/theme-default.css"/>
    <link rel="stylesheet" type="text/css"  href="../static/css/common.css"/>
</head>
<body>
<div class="page-container">
    {{template "pageheader"}}
    <div class="page-content">
        {{template "top"}}
        <ul class="breadcrumb">
            <li><a href="#">Home</a></li>
            <li><a href="#">Tables</a></li>
            <li class="active">Basic</li>
        </ul>
        <div class="page-content-wrap">
            <div class="row">
                <div class="col-md-12">
                    <div class="panel panel-default">
                        <div class="page-title page-category">
                            <h2><span class="fa fa-arrow-circle-o-left"></span> 文章列表</h2>
                        </div>
                        <div class="panel-heading">
                            <div class="input-group push-down-10 col-md-4">
                                <span class="input-group-addon"><span class="fa fa-search"></span></span>
                                <input type="text" class="form-control" placeholder="请输入关键字...">
                                <div class="input-group-btn">
                                    <button class="btn btn-primary">Search</button>
                                </div>
                            </div>
                            <h3 class="panel-title pull-right">
                                <a href="/article/add" class="btn btn-success btn-block btn-func">新增</a>
                            </h3>
                        </div>
                        <div class="panel-body panel-body-table">
                            <div class="table-responsive">
                                <table class="table table-bordered table-striped table-actions">
                                    <thead>
                                    <tr>
                                        <th>编号</th>
                                        <th>名称</th>
                                        <th>分类</th>
                                        <th>浏览量</th>
                                        <th>内容</th>
                                        <th>操作</th>
                                    </tr>
                                    </thead>
                                    <tbody>
                                    {{range .Articles}}
                                    <tr id="{{.Id}}">
                                        <td class="text-center">{{.Id}}</td>
                                        <td><strong>{{.Title}}</strong></td>
                                        <td><strong>{{.Category.Id}}</strong></td>
                                        <td><strong>{{.Views}}</strong></td>
                                        <td><strong>{{html2str .Content}}</strong></td>
                                        <td>
                                            <a href="/article/edit?id={{.Id}}" class="btn btn-default btn-rounded btn-sm">
                                                <span class="fa fa-pencil"></span>
                                            </a>
                                            <a href="/article/show?id={{.Id}}" alt="详情" class="btn btn-default btn-rounded btn-sm">
                                                <i class="fa fa-info"></i>
                                            </a>
                                            <button class="btn btn-danger btn-rounded btn-sm" onClick="delete_row('{{.Id}}');">
                                                <span class="fa fa-times"></span>
                                            </button>
                                        </td>
                                    </tr>
                                    {{end}}
                                    </tbody>
                                </table>
                            </div>

                        </div>
                        <div class="panel-footer">
                          
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
 {{template "footer"}}
<div class="message-box animated fadeIn" data-sound="alert" id="mb-remove-row">
    <div class="mb-container">
        <div class="mb-middle">
            <div class="mb-title"><span class="fa fa-times"></span> Remove <strong>Data</strong> ?</div>
            <div class="mb-content">
                <p>Are you sure you want to remove this row?</p>
                <p>Press Yes if you sure.</p>
            </div>
            <div class="mb-footer">
                <div class="pull-right">
                    <button class="btn btn-success btn-lg mb-control-yes">Yes</button>
                    <button class="btn btn-default btn-lg mb-control-close">No</button>
                </div>
            </div>
        </div>
    </div>
</div>
<!-- END MESSAGE BOX-->
<script>
function delete_row(row){
        var box = $("#mb-remove-row");
        box.addClass("open");
        box.find(".mb-control-yes").on("click",function(){
            jQuery.ajax({
                url: "/category/delete",
                data:{"id":row},
                dataType:"json",
                type:"get",
            }).done(function (data) {
                //layer.msg('提交成功！');      
                box.removeClass("open");
                $("#"+row).hide("slow",function(){
                    $(this).remove();
                });
            }).fail(function(data) { 
               //layer.msg('提交失败!');       
           }); 
        });
        box.find(".mb-control-close").on("click",function(){
            box.removeClass("open");                                                                                                             
        });

}
</script>
</body>
</html>
