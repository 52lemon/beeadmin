<!DOCTYPE html>
<html lang="en">
<head>
    <!-- META SECTION -->
    <title>{{.Article.Title}}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />

    <link rel="icon" href="favicon.ico" type="image/x-icon" />
    <!-- END META SECTION -->

    <!-- CSS INCLUDE -->
    <link rel="stylesheet" type="text/css" id="theme" href="../static/joli/css/theme-default.css"/>
    <!-- EOF CSS INCLUDE -->
    <style>
        .img-thumbnail{width:120px;height:120px;}
    </style>
</head>
<body>
<!-- START PAGE CONTAINER -->
<div class="page-container">

    <!-- START PAGE SIDEBAR -->
        {{template "pageheader"}}
    <!-- END PAGE SIDEBAR -->

    <!-- PAGE CONTENT -->
    <div class="page-content">

        <!-- START X-NAVIGATION VERTICAL -->
        {{template "top"}}
        <!-- END X-NAVIGATION VERTICAL -->

        <!-- START BREADCRUMB -->
        <ul class="breadcrumb">
            <li><a href="#">Home</a></li>
            <li><a href="#">Forms Stuff</a></li>
            <li><a href="#">Form Layout</a></li>
            <li class="active">One Column</li>
        </ul>
        <!-- END BREADCRUMB -->

        <!-- PAGE CONTENT WRAPPER -->
        <div class="page-content-wrap">

            <div class="row">
                <div class="col-md-12">

                    <form class="form-horizontal" id="jvalidate" role="form" method="post" action="/article/modify" enctype="multipart/form-data">
                        <div class="panel panel-default">
                            <div class="panel-heading">
                                <h3 class="panel-title"><strong>新增文章</strong></h3>
                            </div>
                            <div class="panel-body">
                                <input type="hidden" name="tid" value="{{.Article.Id}}" class="form-control"/>
                                <div class="form-group">
                                    <label class="col-md-3 col-xs-12 control-label">标题：</label>
                                    <div class="col-md-6 col-xs-12">
                                        <div>{{.Article.Title}}</div>
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-md-3 col-xs-12 control-label">分类：</label>
                                    <div class="col-md-6 col-xs-12">
                                        <div>{{.Category.id}}</div>
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-md-3 col-xs-12 control-label">缩略图：</label>
                                    <div class="col-md-6 col-xs-12">
                                        <img src="{{makeUrl `80095618-a274-48d9-b117-f23e3e51bc70.jpeg`}}"  class="img-thumbnail"/>
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-md-3 col-xs-12 control-label">简介:</label>
                                    <div class="col-md-6 col-xs-12">
                                        <div class="">{{.Article.Summary}}</div>
                                    </div>
                                </div>

                                <div class="form-group">
                                    <label class="col-md-3 col-xs-12 control-label">内容：</label>
                                    <div class="col-md-6 col-xs-12">
                                        <div class="">{{html2str .Article.Content}}</div>
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-md-3 col-xs-12 control-label">推荐：</label>
                                    <div class="col-md-6 col-xs-12">
                                        <label class="check">
                                            <input type="checkbox" class="icheckbox" checked="checked" disabled/> Checkbox title
                                        </label>
                                    </div>
                                </div>

                            </div>
                            <div class="panel-footer">
                                <button class="btn btn-default">重置</button>
                                <button class="btn btn-primary pull-right">保存</button>
                            </div>
                        </div>
                    </form>

                </div>
            </div>

        </div>
        <!-- END PAGE CONTENT WRAPPER -->
    </div>
    <!-- END PAGE CONTENT -->
</div>
<!-- END PAGE CONTAINER -->

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
<!-- END MESSAGE BOX-->

<!-- START PRELOADS -->
<audio id="audio-alert" src="../static/joli/audio/alert.mp3" preload="auto"></audio>
<audio id="audio-fail" src="../static/joli/audio/fail.mp3" preload="auto"></audio>
<!-- END PRELOADS -->

<!-- START SCRIPTS -->
<!-- START PLUGINS -->
<script type="text/javascript" src="../static/joli/js/plugins/jquery/jquery.min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/jquery/jquery-ui.min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/bootstrap/bootstrap.min.js"></script>
<!-- END PLUGINS -->

<!-- THIS PAGE PLUGINS -->
<script type='text/javascript' src='../static/joli/js/plugins/icheck/icheck.min.js'></script>
<script type="text/javascript" src="../static/joli/js/plugins/mcustomscrollbar/jquery.mCustomScrollbar.min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/bootstrap/bootstrap-file-input.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/bootstrap/bootstrap-select.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/tagsinput/jquery.tagsinput.min.js"></script>
<!-- END THIS PAGE PLUGINS -->
<script type="text/javascript" src="../static/joli/js/plugins/codemirror/codemirror.js"></script>
<script type='text/javascript' src="../static/joli/js/plugins/codemirror/mode/htmlmixed/htmlmixed.js"></script>
<script type='text/javascript' src="../static/joli/js/plugins/codemirror/mode/xml/xml.js"></script>
<script type='text/javascript' src="../static/joli/js/plugins/codemirror/mode/javascript/javascript.js"></script>
<script type='text/javascript' src="../static/joli/js/plugins/codemirror/mode/css/css.js"></script>
<script type='text/javascript' src="../static/joli/js/plugins/codemirror/mode/clike/clike.js"></script>

<script type="text/javascript" src="../static/joli/js/plugins/summernote/summernote.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/dropzone/dropzone.min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/fileinput/fileinput.min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/filetree/jqueryFileTree.js"></script>
<!-- START TEMPLATE -->
<script type="text/javascript" src="../static/joli/js/settings.js"></script>

<script type="text/javascript" src="../static/joli/js/plugins.js"></script>
<script type="text/javascript" src="../static/joli/js/actions.js"></script>
<!-- END TEMPLATE -->
<script>
    var editor = CodeMirror.fromTextArea(document.getElementById("codeEditor"), {
        lineNumbers: true,
        matchBrackets: true,
        mode: "application/x-httpd-php",
        indentUnit: 4,
        indentWithTabs: true,
        enterMode: "keep",
        tabMode: "shift"
    });
    editor.setSize('100%','420px');
    $(function(){
        $("#file-simple").fileinput({
            showUpload: false,
            showCaption: false,
            browseClass: "btn btn-danger",
            fileType: "any"
        });
    });
    var jvalidate = $("#jvalidate").validate({
        ignore: [],
        rules: {
            login: {
                required: true,
                minlength: 2,
                maxlength: 8
            },
            password: {
                required: true,
                minlength: 5,
                maxlength: 10
            },
            're-password': {
                required: true,
                minlength: 5,
                maxlength: 10,
                equalTo: "#password2"
            },
            age: {
                required: true,
                min: 18,
                max: 100
            },
            email: {
                required: true,
                email: true
            },
            date: {
                required: true,
                date: true
            },
            credit: {
                required: true,
                creditcard: true
            },
            site: {
                required: true,
                url: true
            }
        }
    });
</script>
<!-- END SCRIPTS -->
</body>
</html>






