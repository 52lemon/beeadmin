<!DOCTYPE html>
<html lang="en">
<head>
    <title>分类列表</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <link rel="icon" href="favicon.ico" type="image/x-icon" />
    <link rel="stylesheet" type="text/css" id="theme" href="../static/joli/css/theme-default.css"/>
    <link rel="stylesheet" type="text/css"  href="../static/css/common.css"/>
    <style>
      .nav-tabs, .nav-tabs.nav-justified {padding: 0px;}
      .nav-tabs{background-color:#fff;border-bottom:1px solid #ddd;}
      .nav-tabs>li{margin-bottom:-2px;}
      .nav-tabs>li.active{border-left:1px solid #ccc;border-right:1px solid #ccc;margin-bottom:-2px;}
      .nav-tabs>li>a{margin-right:0!important;background-color:#fff;}
      .nav-tabs>li.active>a{margin-right:0!important;}
    </style>
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
                    <div class="panel panel-default tabs">                            
                        <ul class="nav nav-tabs" role="tablist">
                            <li class="active">
                                <a href="#tab-first" role="tab" data-toggle="tab">分类列表</a>
                            </li>
                            <li>
                                <a href="#tab-second" role="tab" data-toggle="tab">分类树</a>
                            </li>
                        </ul>
                        <div class="panel-body tab-content">
                            <div class="tab-pane active" id="tab-first">
                            <div class="table-responsive">
                                <table class="table table-bordered table-striped table-actions">
                                    <thead>
                                    <tr>
                                        <th>编号</th>
                                        <th>名称</th>
                                        <th>父类</th>
                                        <th>备注</th>
                                        <th>操作</th>
                                    </tr>
                                    </thead>
                                    <tbody>
                                    {{range .Categories}}
                                    <tr id="{{.Id}}">
                                        <td class="text-center">{{.Id}}</td>
                                        <td><strong>{{.Title}}</strong></td>
                                        <td><strong>{{.Parent.Id}}</strong></td>
                                        <td><strong>{{.Description}}</strong></td>
                                        <td>
                                            <a href="/category/edit?id={{.Id}}" class="btn btn-default btn-rounded btn-sm">
                                                <span class="fa fa-pencil"></span>
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
                            <div class="tab-pane" id="tab-second">
                                
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
<!-- MESSAGE BOX-->
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
