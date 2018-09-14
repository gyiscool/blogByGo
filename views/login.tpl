<!DOCTYPE html>
<html lang="zh_cn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>行车聘-人才超市</title>

    <link href="../../fonts/iconfont.css" rel="stylesheet">
    <link href="../../styles/reset.css" rel="stylesheet">
    
    <script src="../../scripts/jquery.min.js"></script>

    <script src="login.js"></script>
    <link href="login.css" rel="stylesheet">
</head>
<body>
    <div class="container tc mask">
        <div class="login-box pa">
            <div class="login-head usn">
                <!-- <img src="../../images/logo-academy.png" alt="logo"> -->
                <i class="iconfont icon-logo"></i>
            </div>
            <!-- 用户名密码登录表单 -->
            <form class="login-body login-pwd dn" action="/user/login_go" method="post">
                <label class="pr">
                    <i class="iconfont icon-phone pa"></i>
                    <input type="text" name="account" placeholder="手机号" class="mgb40">
                </label>
                <label class="pr">
                    <i class="iconfont icon-password pa"></i>
                    <input type="password" name="passwd" placeholder="密码">
                    <i class="iconfont icon-show icon-hide pa"></i>
                </label>
                <div class="msg">用户名或密码错误</div>
                <div class="opt-box clearfix">
                    <div class="fl">
                        <input type="checkbox" name="nextpass" id="nextpass" class="dn"><label for="nextpass" class="usn"><i class="iconfont icon-uncheck"></i>下次登录记住密码</label>
                        <a href="#code" class="abtn">短信验证码登录</a>
                    </div>
                    <button type="submit" class="fr">登录</button>
                </div>
            </form>

            <form class="login-body login-code dn">
                    <label class="pr">
                        <i class="iconfont icon-phone pa"></i>
                        <input type="text" name="account" placeholder="手机号"  class="mgb40">
                    </label>
                    <label class="pr">
                        <i class="iconfont icon-verifycode pa"></i>
                        <input type="text" name="vcode" placeholder="验证码">
                        <a href="javascript:;" class="pa abtn sendcodebtn" id="sendcode2login">发送短信验证码</a>
                    </label>
                    <div class="msg">用户名或验证码错误</div>
                    <div class="opt-box clearfix">
                        <a href="#pwd" class="fl abtn">帐号密码登录</a>
                        <button type="submit" class="fr">登录</button>
                    </div>
            </form>

            <div class="from-bottom login-bottom pr dn">
                <span class="tips pa">还没有帐号？</span>
                <a href="#company" class="icon-btn">
                    <i class="iconfont icon-company"></i>
                    <span>我是企业</span>
                </a>
                <a href="#school" class="icon-btn">
                    <i class="iconfont icon-school"></i>
                    <span>我是院校</span>
                </a>
            </div>
            
            <!-- 注册表单 -->
            <form class="login-body login-register dn">
                <div class="register-top">
                    <input type="radio" name="user_type" value="1" class="dn" id="companyType"><label for="companyType">企业</label>
                    <input type="radio" name="user_type" value="2" class="dn" id="schoolType"><label for="schoolType">院校</label>
                </div>

                <label class="pr">
                    <i class="iconfont icon-phone pa"></i>
                    <input type="text" name="account" placeholder="手机号"  class="mgb40">
                </label>
                <label class="pr">
                    <i class="iconfont icon-verifycode pa"></i>
                    <input type="text" name="vcode" placeholder="验证码">
                    <a href="javascript:;" class="pa abtn sendcodebtn" id="sendcode2regiter">发送短信验证码</a>
                </label>
                <div class="msg">用户名或验证码错误</div>

                <div class="opt-box">
                    <button type="submit">注册</button>
                    <p>注册代表你同意<a href="#terms">《人才超市用户协议》</a></p>
                </div>
                
                <div class="from-bottom pr">
                    <span class="tips pa">已有帐号？</span>
                    <a href="#code" id="directLogin">直接登录 &gt;</a>
                </div>
            </form>

        </div>
    </div>
</body>
</html>