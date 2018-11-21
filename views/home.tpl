<!DOCTYPE html>

<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=10,IE=9,IE=8">
	<meta name="baidu-site-verification" content="2m6jxCb2Ab" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=0, minimum-scale=1.0, maximum-scale=1.0">
	<meta name="keywords" content="java,spring,redis,mysql,go,php,java,beego,yii2,设计模式,nginx,响应时间,性能优化">
	<title>耿阳的个人博客</title>
<script>
window._deel = {name: '耿阳的个人博客',url: '8', ajaxpager: '/ss/ss', commenton: 0, roll: [0,0]}
</script>
	<link rel="dns-prefetch" href="http://libs.baidu.com/">
	<link rel="dns-prefetch" href="http://s.w.org/">
	<link rel="shortcut icon" href="/static/images/faviconq.ico" type="image/x-icon" />
	<!--<link rel="stylesheet" id="style-css" href="/static/css/style/reset.css" type="text/css" media="all"> -->
	<link rel="stylesheet" id="style-css" href="/static/css/style/style.css" type="text/css" media="all">
	<link rel="stylesheet" id="style-cssa" href="/static/css/home/style.css" type="text/css" media="all">
	<script type="text/javascript" src="/static/js/home/jquery.min.js"></script>
	<script type="text/javascript" src="/static/js/home/jquery.js"></script>
	<script src="/static/js/home/wp-emoji-release.min.js" type="text/javascript" defer=""></script>
	<script type="text/javascript" src="/static/js/home/wp-embed.min.js"></script>
	<script type="text/javascript" src="/static/js/home/page.js"></script>
	<meta name="description" content="耿阳的个人博客">
	<!--分页看看吧-->
	
</headpageUrl>
<body class="home blog logged-in" url="{{.pageUrl}}">

	<header id="header" class="header">
	<div class="container-inner">
		<div class="yusi-logo">
			<a href="#">
				<h1>
					<span class="yusi-mono">耿阳</span>
					<span class="yusi-bloger">de个人博客</span>
				</h1>
			</a>
		</div>
	</div>

		<div id="nav-header" class="navbar">
			<ul class="nav">
				<li id="menu-item-3307" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-3307">
						<a href="/">首页</a>
				</li>
				{{range $index, $elem := .terms}}
					<li id="menu-item-3307" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-3307">
						<a href="/term/{{$elem.Uid}}/articles">{{$elem.Title}}</a>
						{{if .Terms}}
							<ul class="sub-menu">
								{{range $index1, $elem1 := .Terms}}
									<li id="menu-item-3301" class="menu-item menu-item-type-taxonomy menu-item-object-category menu-item-3301">
										<a href="/term/{{$elem1.Uid}}/articles">{{$elem1.Title}}</a>
									</li>
								{{end}}
							</ul>
						{{else}}
							
						{{end}}
					</li>
				{{end}}
				<li style="float:right;">
					<div class="toggle-search">
						<i class="fa fa-search"></i>
					</div>
					<div class="search-expand" style="display: none;">
						<div class="search-expand-inner">
							<form method="get" class="searchform themeform" onsubmit="location.href='?title=' + encodeURIComponent(this.s.value).replace(/%20/g, '+'); return false;" action="">
								<div> 
									<input type="ext" class="search" name="s" onblur="if(this.value=='')this.value='search...';" onfocus="if(this.value=='search...')this.value='';" value="search...">
								</div>
							</form>
						</div>
					</div>
				</li>
			</ul>
			<div class="ias_trigger" style="display:none;"></div>
			<div class="screen-mini"><button data-type="screen-nav" class="btn btn-inverse screen-nav">
				<i class="fa fa-list"></i></button>
			</div>
		</div>
		
	</header>

	<section class="container">
		<div class="speedbar">
			<div class="toptip"><strong class="text-success"><i class="fa fa-volume-up"></i> 快上车，来不及解释了</strong> </div>
		</div>

		<!--列表页-->
		<div class="content-wrap">
			<div class="content">
			
				{{range $index, $elem := .articles}}
					<article class="excerpt">
						<header>
							<a class="label label-important" href="/article/{{$elem.Uid}}">{{$elem.Term.Title}}<i class="label-arrow"></i>
							</a>
							<h2><a target="_blank" href="/article/{{$elem.Uid}}" title="11月5日">{{$elem.Title}}  </a></h2>
						</header>
						<div class="focus">
						</div>
						<!--<span class="note"> {{$elem.Brief}}</span>-->
						<p class="auth-span">
							<span class="muted">
								<i class="fa fa-user"></i> 
								<a href="javascript:;">{{if $elem.Admin}}{{$elem.Admin.Nick_name}}{{end}}</a>
							</span>
							<span class="muted"><i class="fa fa-clock-o"></i> {{$elem.Cdate}}</span>
							<span class="muted"><i class="fa fa-eye"></i> {{$elem.Views}}℃</span>
							<span class="muted">
								<i class="fa fa-comments-o"></i>
								<a target="_blank" href="/article/{{$elem.Uid}}#comments">{{$elem.Comments}}评论</a>
							</span>
							<span class="muted">
								<a href="javascript:;" data-action="like" data-id="{{$elem.Uid}}" id="Addlike" class="action"><i class="fa fa-heart-o"></i><span class="count">{{$elem.Zans}}</span>喜欢</a>
							</span>
						</p>
					</article>
				{{end}}

					<div class="pagination" style="{{if .string}} display:block;{{end}}">
						<ul>
						{{str2html .string}}
						</ul>
					</div>

			</div>
		</div>
		
		<!--侧边栏-->
		<aside class="sidebar">	
			<div class="widget widget_text">
				<div class="textwidget">
					</div>
				</div>
			</div>
			<div class="widget widget_text">
				<div class="title">
					<h2><i class="themify-menu-icon ti-themify-favicon"></i> 写在前头</h2>
				</div>			
				<div class="textwidget" style="padding: 5px 20px 20px 20px;">
					<div class="alert alert-info">
						<p>这是我的个人博客，如果有问题需要讨论，请评论！或者发送邮件643073032@qq.com</p>
						<!--<p>
							<a href="http://weibo.com/live" target="_blank" data-slimstat="5"><i class="fa fa-send"></i> gy</a>
						</p>-->
					</div>
				</div>
			</div>

			<div class="widget d_postlist">
			  <div class="title">
			    <h2>最新发布</h2></div>
			  <ul>
			{{range $index, $elem := .newArticles}}
			    <li>
			      <a href="/article/{{$elem.Uid}}" title="{{$elem.Title}}" style='padding-left:0px;'>
			        <span class="text" style="height:30px;margin-left:10px">{{$elem.Title}}</span>
			        <span class="muted">{{$elem.Cdate}}</span>
			        <span class="muted" style="float: right;">{{$elem.Comments}}评论</span></a>
			    </li>
				{{end}}			
			  </ul>
			</div>
		</aside>
				
	</section>
	<footer class="footer">
		<div class="footer-inner">
			<div class="copyright pull-left">
			 <a href="http://www.gyiscool.com/" title="gy博客">耿阳的个人博客</a> 版权所有，保留一切权利 · <a href="" title="站点地图">站点地图</a>   ·   基于goweb构建   © 2011-2019  ·   托管于腾讯云  ·  <a rel="nofollow" target="_blank" href="#">苏ICP备17004846号-1</a> &amp; <a rel="nofollow" target="_blank" href="http://gyiscool.com">个人存储</a>
			</div>
			<div class="trackcode pull-right"></div>
		</div>
	</footer>
	<div class="rollto">
		<button class="btn btn-inverse" data-type="totop" title="回顶部"><i class="fa fa-arrow-up"></i></button>
	</div>
</body>
<div id="cVim-status-bar" style="top: 0px;"></div>
</html>
