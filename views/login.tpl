<!DOCTYPE html>
<html lang="en" class="body-full-height">
<head>
    <!-- META SECTION -->
    <title>登陆</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <link rel="icon" href="favicon.ico" type="image/x-icon" />
    <!-- END META SECTION -->

    <!-- CSS INCLUDE -->
    <link rel="stylesheet" type="text/css" id="theme" href="../static/joli/css/theme-blue.css"/>
    <!-- EOF CSS INCLUDE -->
</head>
<body>

<div class="login-container">

    <div class="login-box animated fadeInDown">
        <div class="login-logo"></div>
        <div class="login-body">
            <div class="login-title"><strong>Welcome</strong>, Please login</div>
            <form action="/" class="form-horizontal" method="post" id="iform">
                <div class="form-group">
                    <div class="col-md-12">
                        <input type="text" id="uname" name="uname" class="form-control" placeholder="账号"/>
                    </div>
                </div>
                <div class="form-group">
                    <div class="col-md-12">
                        <input type="password" id="pwd" name="pwd" class="form-control" placeholder="密码"/>
                    </div>
                </div>
                <div class="form-group">
                    <div class="col-md-6">
                        <a href="#" class="btn btn-link btn-block">忘记密码?</a>
                    </div>
                    <div class="col-md-6">
                        <button type="button" class="btn btn-info btn-block"  onclick="return checkInput();">登陆</button>
                    </div>
                </div>
            </form>
        </div>
        <div class="login-footer">
            <div class="pull-left">
                &copy; 2014 AppName
            </div>
            <div class="pull-right">
                <a href="#">About</a> |
                <a href="#">Privacy</a> |
                <a href="#">Contact Us</a>
            </div>
        </div>
    </div>

</div>
<script type="text/javascript">
    var form = document.getElementById("iform");
    function checkInput() {
        var uname = document.getElementById("uname");
        var pwd = document.getElementById("pwd");
        if (uname.value.length == 0) {
            alert("请输入管理员帐号");
            return false;
        }

        if (pwd.value.length == 0) {
            alert("请输入管理员密码");
            return false;
        }
        form.submit();
    }
</script>
</body>
</html>






