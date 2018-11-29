<!--列表页-->
<div class="content-wrap">
	<div class="content">
	
		<!--展示分类-->
		{{if .termTitle}}
			<header class="archive-header" style="margin-bottom: 10px"> 
				<h1>
					<i class="fa fa-folder-open"></i>  &nbsp;分类：{{.termTitle}}  
					</a>
				</h1>
			</header>
		{{end}}
	
		{{range $index, $elem := .articles}}
			<article class="excerpt">
				<header>
					<a class="label label-important" href="/article/{{$elem.Uid}}.html">{{$elem.Term.Title}}<i class="label-arrow"></i>
					</a>
					<h2><a target="_blank" href="/article/{{$elem.Uid}}.html" title="11月5日">{{$elem.Title}}  </a></h2>
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
						<a target="_blank" href="/article/{{$elem.Uid}}.html#comments">{{$elem.Comments}}评论</a>
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

