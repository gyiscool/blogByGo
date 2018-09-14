<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>管理系统</title>
  <link rel="stylesheet" href="/static/layui/src/css/layui.css">
</head>
<body class="layui-layout-body">
<div class="layui-layout layui-layout-admin">
  <div class="layui-header">
    <div class="layui-logo">管理系统</div>
    <ul class="layui-nav layui-layout-right">
      <li class="layui-nav-item">
        <a href="javascript:;">
          <img src="http://t.cn/RCzsdCq" class="layui-nav-img">
          gy
        </a>
        <dl class="layui-nav-child">
          <dd><a href="">基本资料</a></dd>
          <dd><a href="">安全设置</a></dd>
        </dl>
      </li>
      <li class="layui-nav-item"><a href="">退了</a></li>
    </ul>
  </div>
  
  <div class="layui-side layui-bg-black">
    <div class="layui-side-scroll">
      <!-- 左侧导航区域（可配合layui已有的垂直导航） -->
   <ul class="layui-nav layui-nav-tree" id="L_demoNav" lay-filter="test">
        <li class="layui-nav-item layui-nav-itemed">
          <a href="javascript:;">文章管理<span class="layui-nav-more"></span></a>
          <dl class="layui-nav-child">
            <dd lay-id="articles"><a href="/adm/articles">列表</a></dd>
            <dd lay-id="article"><a  href="/adm/article">新增</a></dd>
			<dd lay-id="comments"><a  href="/adm/comments">回复列表</a></dd>
          </dl>
        </li>
        <li class="layui-nav-item layui-nav-itemed">
          <a href="javascript:;">用户管理<span class="layui-nav-more"></span></a>
          <dl class="layui-nav-child">
            <dd lay-id="users"><a href="/adm/users">列表</a></dd>
            <!--<dd lay-id="user"><a href="javascript:;">新增</a></dd>-->
          </dl>
        </li>
        <li class="layui-nav-item">
		 <dd lay-id="info"><a href="javascript:;">信息</a></dd>
        </li>
        <li class="layui-nav-item"><a href="">大数据</a></li>
        <li class="layui-nav-item layui-hide" id="L_demoNavReset"><a href="javascript:;" onclick="L_demoNav.className='layui-nav layui-nav-tree';L_demoNavReset.className='layui-nav-item layui-hide';">还原该导航演示</a></li>
      <span class="layui-nav-bar" style="top: 22.5px; height: 0px; opacity: 0;"></span></ul>
    </div>
  </div>

  <!--<div class="layui-body">-->
    <!-- 内容主体区域 -->
    <!--<div style="padding: 15px;">内容主体区域</div> -->
	{{.LayoutContent}}
  <!--</div>-->
  
  <div class="layui-footer">
    <!-- 底部固定区域 -->
    © gy
  </div>
</div>
<script>
//动态选中某个元素 
layui.use('element', function(){
	var element = layui.element; //导航的hover效果、二级菜单等功能，需要依赖element模块
  	var $ = layui.jquery;
	var layid = location.pathname.replace(/^\/adm\//, '');
	$("[lay-id='"+layid+"']").addClass("layui-this")
	
	//监听导航点击
	/*element.on('nav(test)', function(data){
		location.hash = 'test='+ this.getAttribute('lay-id');
  });*/
});
</script>

</body>
</html>