<!DOCTYPE html>
<html lang="en">
<head>
    <title>beeadmin首页</title>
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
     </div>
</div>
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

<!-- START THIS PAGE PLUGINS-->
<script type='text/javascript' src='../static/joli/js/plugins/icheck/icheck.min.js'></script>
<script type="text/javascript" src="../static/joli/js/plugins/mcustomscrollbar/jquery.mCustomScrollbar.min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/scrolltotop/scrolltopcontrol.js"></script>

<script type="text/javascript" src="../static/joli/js/plugins/morris/raphael-min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/morris/morris.min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/rickshaw/d3.v3.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/rickshaw/rickshaw.min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/jvectormap/jquery-jvectormap-1.2.2.min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/jvectormap/jquery-jvectormap-world-mill-en.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/bootstrap/bootstrap-datepicker.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/owl/owl.carousel.min.js"></script>

<script type="text/javascript" src="../static/joli/js/plugins/moment.min.js"></script>
<script type="text/javascript" src="../static/joli/js/plugins/daterangepicker/daterangepicker.js"></script>
<!-- END THIS PAGE PLUGINS-->

<!-- START TEMPLATE -->
<script type="text/javascript" src="../static/joli/js/settings.js"></script>

<script type="text/javascript" src="../static/joli/js/plugins.js"></script>
<script type="text/javascript" src="../static/joli/js/actions.js"></script>

<script type="text/javascript" src="../static/joli/js/demo_dashboard.js"></script>
<!-- END TEMPLATE -->
<!-- END SCRIPTS -->
</body>
</html>






