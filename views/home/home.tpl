<!--列表页-->
<div class="content-wrap">
	<div class="content">
	
		<!--展示分类-->
		<header class="archive-header" style="margin-bottom: 10px"> 
			<h1>
				<i class="fa fa-folder-open"></i>  &nbsp;分类：{{.termTitle}}  
				</a>
			</h1>
		</header>
	
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
					<!--<span class="muted">
						<i class="fa fa-user"></i> 
						<a href="javascript:;">{{if $elem.Admin}}{{$elem.Admin.Nick_name}}{{end}}</a>
					</span>
					-->
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