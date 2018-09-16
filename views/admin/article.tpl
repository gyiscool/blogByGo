<div class="layui-body layui-tab-content site-demo site-demo-body">
<fieldset class="layui-elem-field layui-field-title" style="margin-top: 20px;">
  <legend>文章编辑</legend>
</fieldset>
	<form class="layui-form" action="">
	<input type="hidden" name="Uid" value="{{if .article}}{{.article.Uid}}{{end}}" lay-skin="switch">
  <div class="layui-form-item">
    <label class="layui-form-label">输入框</label>
    <div class="layui-input-block">
      <input type="text" name="Title" required value="{{if .article}}{{.article.Title}}{{end}}" lay-verify="required" placeholder="请输入标题" autocomplete="off" class="layui-input">
    </div>
  </div>
  
  <div class="layui-form-item">
    <label class="layui-form-label">分类</label>
    <div class="layui-input-block">
	{{if .article}}      
		<select name="Term_id" lay-verify="required">
			<option value="0">无</option>
			{{$m := $.article.Term.Uid }}
			{{range $index, $elem := .terms}}
				<option value="{{$elem.Uid}}" {{if eq ($elem.Uid)  $m }}selected{{end}} >{{$elem.Title}}</option>
					{{if .Terms}}
						{{range $index1, $elem1 := .Terms}}
							<option value="{{$elem1.Uid}}" {{if eq ($elem1.Uid)  $m }}selected{{end}} >&nbsp;&nbsp;&nbsp;&nbsp;{{$elem1.Title}}</option>
						{{end}}
					{{end}}
			{{end}}
	
		</select>
	{{else}}
		<select name="Term_id" lay-verify="required">
			<option value="0">无</option>
			{{range $index, $elem := .terms}}
				<option value="{{$elem.Uid}}" >{{$elem.Title}}</option>
					{{if .Terms}}
						{{range $index1, $elem1 := .Terms}}
							<option value="{{$elem1.Uid}}" >&nbsp;&nbsp;&nbsp;&nbsp;{{$elem1.Title}}</option>
						{{end}}
					{{end}}
			{{end}}
	
		</select>

	
	{{end}}
    </div>
  </div>

  <div class="layui-form-item">
    <label class="layui-form-label">是否隐藏</label>
    <div class="layui-input-block">
      <input type="checkbox" name="switch" lay-skin="switch">
    </div>
  </div>

  <div class="layui-form-item layui-form-text">
    <label class="layui-form-label">简介</label>
    <div class="layui-input-block">
      <textarea placeholder="请输入内容" name="Brief" class="layui-textarea">{{if .article}}{{.article.Brief}}{{end}}</textarea>
    </div>
  </div>

 
  <div class="layui-form-item layui-form-text">
    <label class="layui-form-label">内容</label>
    <div class="layui-input-block">
      <textarea name="Content" id="Content" placeholder="请输入内容" class="layui-textarea">{{if .article}}{{.article.Content}}{{end}}</textarea>
    </div>
  </div>
  <div class="layui-form-item">
    <div class="layui-input-block">
      <button class="layui-btn"  lay-submit data-type="content" lay-filter="formDemo">立即提交</button>
      <button type="reset" class="layui-btn layui-btn-primary">重置</button>
    </div>
  </div>
</form>
</div>
<script src="/static/layui/src/layui.js"></script>
<script>
//获取对应的 分类
var layedit;
var index;
//Demo
layui.use('form', function(){
	
  	var form = layui.form;
  //监听提交
  form.on('submit(formDemo)', function(data){
	console.log(JSON.stringify(data.field));
	data.field.Content = layedit.getContent(index)
	layui.$.ajax({
            type:'POST',
            url :location.pathname,
            data:data.field,
            success:function(res){
                if(res.Code  != 1){
                    layer.msg(res.Message,{icon: 5, shift: 6} );
                }else{
                    layer.msg(res.Message,{icon: 1, shift: -1,time:700},function() {
                        window.location.reload();
                    } );
                }
                return false;
            },
			error:function(res){
				layer.msg("网络错误")
				return false;
			}
        });
    return false;
  });


    form.verify({
        username: function(value, item){ //value：表单的值、item：表单的DOM对象
        }
        //我们既支持上述函数式的方式，也支持下述数组的形式
        //数组的两个值分别代表：[正则匹配、匹配不符时的提示文字]
        ,pass: [
            /^[\S]{6,18}$/
            ,'密码必须6到12位，且不能出现空格'
        ]
    });


});

layui.use('layedit', function(){
	

	var $ = layui.jquery;
  	layedit = layui.layedit;
	layedit.set({
		uploadImage: {
			url: '/adm/file' //接口url
			,type: 'post' //默认post
		}
	});

	index= layedit.build('Content'); //建立编辑器
	layedit.sync(index)
  	
  //编辑器外部操作
 /* var active = {
    content: function(){
     // alert(layedit.getContent(index)); //获取编辑器内容
	//$("#Content").text(layedit.getContent(index))
    }
    ,text: function(){
      //alert(layedit.getText(index)); //获取编辑器纯文本内容
    }
    ,selection: function(){
      //alert(layedit.getSelection(index));
    }
  };
  
  $('#changess').on('click', function(){
    var type = $(this).data('type');
    active[type] ? active[type].call(this) : '';
  });

*/
});

layui.use('upload', function(){
  var upload = layui.upload;
   var $ = layui.jquery;
  //执行实例
  var uploadInst = upload.render({
    elem: '#test1' //绑定元素
    ,url: '/adm/file' //上传接口
    ,done: function(res){
		$(".layui-upload-img").attr('src',res.Data.FileUrl);
		$("input[name='Headimg']").attr('value',res.Data.FileUrl);
      //上传完毕回调
    }
    ,error: function(){
      //请求异常回调
    }
  });
});

</script>