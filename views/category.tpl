<!DOCTYPE html>
<html lang="en">
<head>
    <title>分类列表</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <link rel="icon" href="favicon.ico" type="image/x-icon" />
    <link rel="stylesheet" type="text/css" id="theme" href="../static/joli/css/theme-default.css"/>
    <link rel="stylesheet" type="text/css" id="theme" href="/static/vakata-jstree-3.3.3/dist/themes/default/style.min.css"/>
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
                                <div class="panel-heading">
                                    <div class="input-group push-down-10 col-md-4">
                                       <span class="input-group-addon"><span class="fa fa-search"></span></span>
                                       <input type="text" class="form-control" placeholder="请输入关键字...">
                                       <div class="input-group-btn">
                                           <button class="btn btn-primary">Search</button>
                                       </div>
                                    </div>
                                    <h3 class="panel-title pull-right">
                                        <a href="/category/add" class="btn btn-success btn-block btn-func">新增</a>
                                    </h3>
                                </div>
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
                                <div id="jstree" class="demo col-xs-5">
                                    <ul>
                                    {{range .Trees}} 
                                        {{if .Haschildren}}
                                            <li id={{.Id}} data-jstree='{ "opened" : true }'>{{.Text}} <ul>
                                        {{else}}
                                            <li id={{.Id}} data-jstree='{"icon":"//jstree.com/tree.png"}'>{{.Text}}</li>
                                        {{end}}
                                        {{range .Hero}}
                                            </ul></li>
                                        {{end}}
                                    {{end}}
                                    </ul>
                                </div> 
                                <div class="col-xs-4" style="margin-top:50px;">
                                    <p><label>分类名称：</label><span id="c_name"></span></p>
                                    <p><label>父类名称：</label><span id="c_parent"></span></p>
                                    <p><label>分类描述：</label><span id="c_description"></span></p>
                                    <form action="/category/edit" method="get" onsubmit="return test();">
                                        <input type="hidden" name="id" id="c_edit" value />
                                        <button type="submit" class="btn btn-info pull-right" id="submits">编辑</button>
                                    </form>
                                </div>
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
<script src="/static/vakata-jstree-3.3.3/dist/jstree.min.js"></script>
<script>
//$('#jstree').jstree();
function test(){
    if ( 1 == jQuery("#c_edit").val() ) {alert("不能编辑系统预置分类"); return false;}
    else if ( jQuery("#c_edit").val() ) { return true;}
    else { alert("你还没选择一个分类"); return false;};
}

function context_menu(node){
    var tree = jQuery('#jstree').jstree(true);
    // The default set of all items
    var items = {
        "Create": {
            "separator_before": false,
            "separator_after": false,
            "label": "新  建",
            "action": function (obj) { 
                var jQuerynode = tree.create_node(node);
                tree.edit(jQuerynode);
            }
        },
        "Rename": {
            "separator_before": false,
            "separator_after": false,
            "label": "重命名",
            "action": function (obj) { 
                tree.edit(node);
            }
        },
        "Remove": {
            "separator_before": true,
            "separator_after": false,
            "label": "删  除",
            "action": function (obj) { 
                if(confirm('确认删除该分类?')){
                    tree.delete_node(node);
                    jQuery.noConflict().ajax({
                        url: "/ctree/delete",
                        data: {"id":node.id},
                        type: "post",
                        dataType: "json",
                        success: function(data,textStatus){
                            //location.reload();      
                        },error: function(XMLHttpRequest, textStatus, errorThrown){
                            alert("删除失败");
                        }
                    });
                }
            }
        }
    };
    return items;
}
var flags='false';
jQuery(function () {
    jQuery('#jstree').bind("loaded.jstree", function (e, data) {
         console.dir(e);
         console.dir(data);
     }).jstree({
        "core" : {
            "check_callback" : true,
        },
        "types" : {
                "default" : {
                    "icon" : "fa fa-folder icon-state-warning icon-lg"
                },
                "file" : {
                    "icon" : "fa fa-file icon-state-warning icon-lg"
                },
                "" : {
                    "icon" : "fa fa-file icon-state-warning icon-lg"
                }
        },
        "plugins" : [ 'contextmenu', "dnd", 'sort', "state" ],
        contextmenu : { items: context_menu }
    });
   // jQuery('#jstree').jstree().open_all();
    jQuery('#jstree').on("changed.jstree",function (c,data) {
        jQuery("#c_description").empty();
        jQuery("#c_name").empty();
        jQuery("#c_parent").empty();
        jQuery("#c_edit").val(null);       
        try{
            var c_id = data.node.id;
            console.log('-----c_id = ------- '+c_id);
            jQuery.noConflict().ajax({
                url: "/ctree/get",
                data: {"id":c_id},
                type: "get",
                dataType: "json",
                success: function(e){
                    jQuery('#c_description').html(e.Description);
                    jQuery("#c_name").html(e.Title);
                    jQuery("#c_parent").html(e.Parent.Id);
                    jQuery("#c_edit").val(e.Id);       
                },
                error: function(){alert("获取分类信息失败");}
            });
        }catch(e){
           console.log(e);
        }
    })
    jQuery('#jstree').jstree().open_all();
    jQuery('#jstree').on('create_node.jstree', function (e, data) {
        flags='true';
        mysets = data.node.id.split("_");
        data.node.li_attr.id = data.node.text;
    });
    jQuery('#jstree').on('rename_node.jstree', function (e, data) {
        if(flags=='true'){ 
            jQuery.noConflict().ajax({
                url: "/ctree/save",
                data: {"name":data.text,"parent":data.node.parent},
                type: "post",
                dataType: "json",
                success: function(data,textStatus){
                    flags="false"; 
                   // location.reload();      
                },error: function(XMLHttpRequest, textStatus, errorThrown){
                    console.log(XMLHttpRequest.status);
                    console.dir(XMLHttpRequest);
                    alert("新增分类失败");
                }
            });
        }else{
            jQuery.noConflict().ajax({
                url: "/ctree/modify",
                data: {"name":data.text,"id":data.node.id},
                type: "post",
                dataType: "json",
                success: function(data){
                    //location.reload();      
                },
                error: function(){alert("修改分类失败");}
            });
        }
    });
    jQuery('button').on('click', function () {
        jQuery('#jstree').jstree(true).select_node('child_node_1');
        jQuery('#jstree').jstree('select_node', 'child_node_1');
        jQuery.jstree.reference('#jstree').select_node('child_node_1');
    });
});
function delete_row(row){
        var box = $("#mb-remove-row");
        box.addClass("open");
        box.find(".mb-control-yes").on("click",function(){
            jQuery.ajax({
                url: "/ctree/delete",
                data:{"id":row},
                dataType:"json",
                type:"post",
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
