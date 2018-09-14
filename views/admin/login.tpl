<!DOCTYPE html>
<html>
<head>
	<title>邻里格</title>
    <script type="Text/Javascript" language="JavaScript">
        if (window.top != window){
           // window.top.location.href = '{:U(login/logout)}';
			window.top.location.href = document.location.href;
			//window.parent.frames.location.href="../welcome.en" 
        }
    </script>
 <meta charset="utf-8">
 <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
  <title> New Document </title>
  <meta name="Generator" content="EditPlus">
  <meta name="Author" content="">
  <meta name="Keywords" content="">
  <meta name="Description" content="">
    <meta http-equiv="pragram" content="no-cache">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>后台登录</title>
    <style>
   		 .homeAdminContent{position: absolute;left:0;right: 0;bottom: 0;top:0;min-width: 1100px;min-height: 600px;}
        .footCopyright{position: absolute;left:0;right: 0;bottom: 0;line-height:35px;color:#fff;color:rgba(255,255,255,0.7);background-color:#333;background-color:rgba(0,0,0,0.5);font-size:12px;text-align: center;}
        .homeAdmin{position: absolute;left:0;right: 0;top: 0;bottom:0;background-size: cover;background-position: center;background-repeat: no-repeat;background-image: url(/static/img/homeAdminBg.jpg);min-width: 1100px;text-align: center;}
        .homeAdmin>span{display: inline-block;height: 100%;vertical-align: middle;}
        .adminLoginForm{display: inline-block;vertical-align: middle;color:#fff;background-color: #fff;background-color: rgba(0,0,0,0.4);width:500px;padding:30px 0 40px 0;border-radius: 10px;}
        .adminlogoHead{width:220px;margin-bottom:17px;}
        .adminForm{display:block;width:330px;margin:0 auto;margin-bottom:22px;border-radius: 5px;border:1px solid #fff;height: 45px;}
        input.adminForm{color:#fff;background-color:#e7cd02;border: none;font-size:16px;font-weight: bold;outline: none;margin-top:50px; }
        input.adminForm:hover{background-color: #CDBC38;color:#333;}
        .adminForm img{float:left;width:28px;margin-top:7px;margin-left:8px;}
        .adminForm input{float:right;height:28px;width:283px;margin-top:7px;background-color: transparent;color:#fff;border: none;outline: none;text-indent: 10px;border-left: 1px solid #fff;font-size:15px;}
        .adminForm select{float:right;height:28px;color:#a9a9a9;width:283px;margin-top:7px;background-color: transparent;border: none;outline: none;text-indent: 10px;border-left: 1px solid #fff;font-size:15px;}		
	</style>
</head>  
<script>



</script>  
<body>
<div class="homeAdminContent">
		<div class="homeAdmin">
			<span></span>
           
			<div class="adminLoginForm">
             <form  autocomplete="on"> 
			<!--	<img class="adminlogoHead" src="__PUBLIC__/am/images/logo.png"/> -->
				<label class="adminForm">
					<img src="/static/img/phone-white.png"/>
					<input type="text" class="form-control"  lay-verify="required"  id="exampleInputText" name="names" placeholder="登陆用户名">
				</label>
				<label class="adminForm">
					<img src="/static/img/adminsuo.png"/>
					<input type="password" class="form-control" name="pwd" id="exampleInputPassword1" placeholder="登陆密码">
				</label>
				<!--<label class="adminForm">
					<img src="/static/img/adminsuo.png"/>
					<select name='' id='test' class="form-control" id="exampleInputText">
						<option value='0'>后台管理员</option>
						<volist name='data' id='vo'>
							<option value='{$vo.building_id}'>{$vo.building_name}</option>
						</volist>
					</select>
				</label>-->
				<input type="submit" class="adminForm" id="adminForm" value="登  录"  href='javascript:;' target="_parent">
				</form>
			</div>
		</div>
   		<div class="footCopyright"><!-- Copyright 2011-2016 我　版权所有　--></div>
    </div>
</body>

<script type="text/javascript" src="/static/js/home/jquery.min.js"></script>
<script src="/static/layui/dist/lay/modules/layer.js"></script>
<script>

    document.onkeydown=keyDownSearch;     
    function keyDownSearch(e) {    
        // 兼容FF和IE和Opera    
        var theEvent = e || window.event;    
        var code = theEvent.keyCode || theEvent.which || theEvent.charCode;    
        if (code == 13) {    
           // getTableListData();//具体处理函数  
			//alert(123123);	
			$("#adminForm").click(); 
            return false;    
        }    
        return true;    
    }  


	$("#adminForm").click(function(){
		//alert($("#test option:selected").val());
		if($('input[name="names"]').val() == '' || $('input[name="pwd"]').val() == ''){
		//	alert('请填写完整');
			layer.msg('请填写完整', {icon: 2, shade: 0 ,time:800});
			return false;
		}
		
		$.ajax({
			url:'',
			type:'post',
			data:{names:$('input[name="names"]').val(),pwd:$('input[name="pwd"]').val(),type:$("#test option:selected").val()},
			success:function(rep){
				
				if(rep.Code == 1){
					window.location = "articles";
				}else{
					layer.msg(rep.Message, {icon: 2, shade: 0 ,time:800});
				}
			},
			error:function(){
				layer.msg("网络链接失败", {icon: 2, shade: 0 ,time:800});
			}
		
		});
		return false;
	});

</script>
</html>