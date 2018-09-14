<div class="layui-body layui-tab-content site-demo site-demo-body">
<fieldset class="layui-elem-field layui-field-title" style="margin-top: 20px;">
  <legend>用户列表</legend>
</fieldset>
	<table id="demo1" lay-filter="test"></table>
	<script type="text/html" id="barDemo">
		<!--<a class="layui-btn layui-btn-primary layui-btn-xs" lay-event="detail">查看</a>-->
		<!--<a class="layui-btn layui-btn-xs" lay-event="edit">编辑</a>-->
		<a class="layui-btn layui-btn-danger layui-btn-xs" lay-event="del">删除</a>
	</script>
</div>
<script src="/static/layui/src/layui.js"></script>
<script>

layui.use('table', function(){
  var table = layui.table;

  //第一个实例
  table.render({
    elem: '#demo1'
    ,height: "full "
    ,url: 'users' //数据接口
    ,page: true //开启分页
	,response: {
  		statusName: 'Code' //数据状态的字段名称，默认：code
		,statusCode: 0 //成功的状态码，默认：0
		,msgName: 'Message' //状态信息的字段名称，默认：msg
		,countName: 'Count' //数据总数的字段名称，默认：count
		,dataName: 'Data' //数据列表的字段名称，默认：data
	}      
   
    ,cols: [[ //表头
      {field: 'Id', title: 'ID',sort: true, fixed: 'left'}
      ,{field: 'Name', title: '用户昵称'}
      ,{field: 'Email', title: '邮箱'}
      ,{field: 'ZanNum', title: '总点赞数',sort: true} 
      ,{field: 'CommentNum', title: '总回复数',sort: true}
      ,{field: 'ViewNum', title: '总浏览量', sort: true}
		,{field: 'Cdate', title: '用户创建时间', sort: true}
	  ,{field: 'right', title: '操作',align:'center', toolbar: '#barDemo'}

    ]]
  });

//监听工具条
  table.on('tool(test)', function(obj){
    var data = obj.data;
    if(obj.event === 'detail'){
		location.href="/user/"+data.Id
      layer.msg('ID：'+ data.Id + ' 的查看操作');
    } else if(obj.event === 'del'){
      layer.confirm('真的删除行么', function(index){
		
		layui.$.ajax({
            type:'Delete',
            url :'user/'+data.Id,
            data:data.field,
            success:function(res){
                if(res.Code  != 1){
                    layer.msg(res.Message,{icon: 5, shift: 6} );
                }else{
                    layer.msg(res.Message,{icon: 1, shift: -1,time:700},function() {
                        //window.location.reload();
                    } );
                }
                return false;
            },
			error:function(res){
				layer.msg("网络错误")
			}
        });
		
		
        obj.del();
        layer.close(index);
      });
    } else if(obj.event === 'edit'){ //后台编辑去
		location.href="article/"+data.Uid
      //layer.alert('编辑行：<br>'+ JSON.stringify(data))
    }
  });
});
</script>