<!DOCTYPE html>
<html lang="en">
<head>
    <title>Joli Admin - Responsive Bootstrap Admin Template</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <link rel="icon" href="favicon.ico" type="image/x-icon" />
    <link rel="stylesheet" type="text/css" id="theme" href="../static/joli/css/theme-default.css"/>
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
        <!-- END BREADCRUMB -->

        <!-- PAGE TITLE -->
        <div class="page-title">
            <h2><span class="fa fa-arrow-circle-o-left"></span> 文章列表</h2>
        </div>
        <!-- END PAGE TITLE -->

        <!-- PAGE CONTENT WRAPPER -->
        <div class="page-content-wrap">

            <!-- START RESPONSIVE TABLES -->
            <div class="row">
                <div class="col-md-12">
                    <div class="panel panel-default">

                        <div class="panel-heading">
                            <h3 class="panel-title"><a href="/article/add" class="btn btn-success btn-block">新增</a></h3>
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
                    </div>

                </div>
            </div>
            <!-- END RESPONSIVE TABLES -->

            <!-- END PAGE CONTENT WRAPPER -->
        </div>
    </div>
    <!-- END PAGE CONTENT -->
</div>
<!-- END PAGE CONTAINER -->

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

<!-- MESSAGE BOX-->
<div class="message-box animated fadeIn" data-sound="alert" id="mb-signout">
    <div class="mb-container">
        <div class="mb-middle">
            <div class="mb-title"><span class="fa fa-sign-out"></span> Log <strong>Out</strong> ?</div>
            <div class="mb-content">
                <p>Are you sure you want to log out?</p>
                <p>Press No if youwant to continue work. Press Yes to logout current user.</p>
            </div>
            <div class="mb-footer">
                <div class="pull-right">
                    <a href="pages-login.html" class="btn btn-success btn-lg">Yes</a>
                    <button class="btn btn-default btn-lg mb-control-close">No</button>
                </div>
            </div>
        </div>
    </div>
</div>
<audio id="audio-alert" src="../static/joli/audio/alert.mp3" preload="auto"></audio>
<audio id="audio-fail" src="../static/joli/audio/fail.mp3" preload="auto"></audio>
<script type="text/javascript" src="../static/joli/js/plugins/jquery/jquery.min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/bootstrap/bootstrap.min.js"></script>
<script type='text/javascript' src='../static/joli/js/plugins/icheck/icheck.min.js'></script>
<script type="text/javascript" src="../static/joli/js/plugins/mcustomscrollbar/jquery.mCustomScrollbar.min.js"></script>
<script type="text/javascript" src="../static/joli/js/settings.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins.js"></script>
<script type="text/javascript" src="../static/joli/js/actions.js"></script>
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
