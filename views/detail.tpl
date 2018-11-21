<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=10,IE=9,IE=8">
	<meta name="baidu-site-verification" content="2m6jxCb2Ab" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=0, minimum-scale=1.0, maximum-scale=1.0">
	<meta name="keywords" content="{{.article.Title}},java,spring,redis,mysql,go,php,java,beego,yii2,设计模式,nginx,响应时间,性能优化">
	<title>{{.article.Title}}</title>
<script>
window._deel = {name: '耿阳的个人博客',url: '8', ajaxpager: '/ss/ss', commenton: 0, roll: [0,0]}
</script>
	<link rel="dns-prefetch" href="http://libs.baidu.com/">
	<link rel="shortcut icon" href="/static/images/faviconq.ico" type="image/x-icon" />
	<link rel="dns-prefetch" href="http://s.w.org/">
	<link rel="stylesheet" id="style-css" href="/static/css/style/style.css" type="text/css" media="all">
	<link rel="stylesheet" id="style-cssa" href="/static/css/home/style.css" type="text/css" media="all">
    <link rel="stylesheet" href="/static\editor.md-master\css\editormd.css" />   

	<script type="text/javascript" src="/static/js/home/jquery.min.js"></script>
	<script type="text/javascript" src="/static/js/home/jquery.js"></script>
	<script src="/static/js/home/wp-emoji-release.min.js" type="text/javascript" defer=""></script>
	<script type="text/javascript" src="/static/js/home/wp-embed.min.js"></script>
	<script type="text/javascript" src="/static/js/home/page.js"></script>
	<style>
	.comt-mailme {
		display: block;
	}
	.popover-content a {
     	color: white; 
	}
	.prettyprint.linenums, pre.prettyprint.linenums  {
		box-shadow: inset 0 0 0 #eee, inset 0 0 0 #33b796;
		
	}
	 pre.prettyprint.linenums ol{
		margin: 0 0 0 0px;
	}
	.linenums li{
	   line-height: 1.6;
	}
	.article-content li:before {
		    
		    content: "";
		   margin: 0; 
     		padding: 0; 
     		width: 0px
 	}

 	ul, menu, dir {
 	    list-style-type: disc;
 	}
 	ul ul, ol ul {
    list-style-type: circle;
	}
	.d_postlist ul{
		 list-style: none;

	}

	.article-contentv2 {
	    font-size: 15px;
	    text-indent: 0px;
	    line-height: 26px;

	    word-break: break-all;
	    word-wrap: break-word;
	    position: relative;
	    padding: 10px 20px 20px 20px;
	    background-color: #fff;
	}
	.children{
		list-style: none;
	}
	</style>
	<meta name="description" content="">
</head>
<body class="home blog logged-in">

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
							<form method="get" class="searchform themeform" onsubmit="location.href='/articles?title=' + encodeURIComponent(this.s.value).replace(/%20/g, '+'); return false;" action="#">
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
		<div class="content-wrap">
		<div class="content">
  <header class="article-header">
    <h1 class="article-title">
      <a href="">{{.article.Title}}</a></h1>
    <div class="meta">
      <span id="mute-category" class="muted">
        <i class="fa fa-list-alt"></i>
        <a href="https://www.gyiscool.com">{{.article.Title}}</a></span>
      <span class="muted">
        <i class="fa fa-user"></i>{{if .article.Admin}}{{.article.Admin.Nick_name}}{{end}}</span>
      <time class="muted">
        <i class="fa fa-clock-o"></i>{{.article.Cdate}}</time>
      <span class="muted" style="display:none;">
        <i class="fa fa-eye"></i>17729℃</span>
      <span class="muted">
        <i class="fa fa-comments-o"></i>
        <a href="https://gyiscool.com">{{.article.Comments}}评论</a></span>
    </div>
  </header>
  <div class="banner banner-post">

    
  </div>
  <article class="article-contentv2">
  	<div class="markdown-body editormd-preview-container editormd-preview-active" previewcontainer="true" style="padding:0;">
  		{{if .article.Content}}{{str2html  .article.Content.Html}}{{end}}
  		
  	</div>
	
    <div class="article-social">
      <a href="javascript:;" data-action="ding" data-id="{{.article.Uid}}" id="Addlike" class="action" data-original-title="" title="">
        <i class="fa fa-heart-o"></i>喜欢 (
        <span class="count">{{.article.Zans}}</span>)</a>
      <span class="or">or</span>
      <span class="action action-share bdsharebuttonbox bdshare-button-style0-24" data-bd-bind="1514884100648">
        <i class="fa fa-share-alt"></i>分享 (
        <span class="bds_count" data-cmd="count" title="累计分享0次">0</span>)
        <div class="action-popover">
          <div class="popover top in">
            <div class="arrow"></div>
            <div class="popover-content">
              <a href="#" class="sinaweibo fa fa-weibo" data-cmd="tsina" title="" data-original-title="分享到新浪微博"></a>
              <a href="#" class="bds_qzone fa fa-star" data-cmd="qzone" title="" data-original-title="分享到QQ空间"></a>
              <a href="#" class="tencentweibo fa fa-tencent-weibo" data-cmd="tqq" title="" data-original-title="分享到腾讯微博"></a>
              <a href="#" class="qq fa fa-qq" data-cmd="sqq" title="" data-original-title="分享到QQ好友"></a>
              <a href="#" class="bds_weixin fa fa-weixin" data-cmd="weixin" title="" data-original-title="分享到微信"></a>
              <a href="#" class="bds_more fa fa-ellipsis-h" data-cmd="more" data-original-title="" title=""></a>
            </div>
          </div>
        </div>
      </span>
    </div>
        <p>转载请注明：
      <a href="" data-original-title="" title="">GY的个人博客</a>»
      <a href="" data-original-title="" title="">{{.article.Title}}</a></p>
    <div class="open-message">
      <i class="fa fa-bullhorn"></i>如果你觉得这篇文章或者我分享的主题对你有帮助，请支持我继续更新分享！
      <a style="float:right;text-indent: 0;" href="javascript:;" title="" target="_blank" data-original-title="捐赠本站">捐赠本站</a></div>
  </article>

  <nav class="article-nav" style="margin-bottom: 10px;">
    <span class="article-nav-prev">
		{{if .preArticle.Uid}}
			<i class="fa fa-angle-double-left"></i>
			<a href="/article/{{.preArticle.Uid}}" rel="prev">
			{{.preArticle.Title}}
			</a>
		{{else}}
			无
		{{end}}
	</span>
    <span class="article-nav-next">
		{{if .nextArticle.Uid}}
		    <a href="/article/{{.nextArticle.Uid}}" rel="next">{{.nextArticle.Title}}</a>
		    <i class="fa fa-angle-double-right"></i>
		{{else}}
			无
		{{end}}
    </span>
  </nav>

  <div id="respond" class="no_webshot">
    <form action="" method="post" id="commentform">
      <div class="comt-title">
        <div class="comt-avatar pull-left">
          <img alt="" src="http://gyiscool.com" srcset="http://gyiscool.com" class="avatar avatar-54 photo" height="54" width="54"></div>
        <div class="comt-author pull-left">发表我的评论</div>
        <a id="cancel-comment-reply-link" class="pull-right" href="javascript:;">取消评论</a></div>
      <div class="comt">
        <div class="comt-box">
          <textarea placeholder="写点什么..." class="input-block-level comt-area" name="comment" id="comment" cols="100%" rows="3" tabindex="1" onkeydown="if(event.ctrlKey&amp;&amp;event.keyCode==13){document.getElementById('submit').click();return false};"></textarea>
          <div class="comt-ctrl">
            <button class="btn btn-primary pull-right" type="submit" name="submit" id="submit" tabindex="5">
              <i class="fa fa-check-square-o"></i>提交评论</button>
            <div class="comt-tips pull-right">
              <input type="hidden" name="comment_post_ID" value="{{.article.Uid}}" id="comment_post_ID">
              <input type="hidden" name="comment_parent" id="comment_parent" value="0">
              
          	</div>
           <!-- <span data-type="comment-insert-smilie" class="muted comt-smilie">
              <i class="fa fa-smile-o"></i>表情</span>-->
            <span class="muted comt-mailme">
              <label for="comment_mail_notify" class="checkbox inline" >
                <input type="checkbox" name="comment_mail_notify" id="comment_mail_notify" value="comment_mail_notify" checked="checked">有人回复时邮件通知我</label>
            </span>
          </div>
        </div>
        <div class="comt-comterinfo" id="comment-author-info" style="display: none;">
          <h4>Hi，您需要填写昵称和邮箱！</h4>
          <ul>
            <li class="form-inline">
              	<label class="hide" for="author">昵称</label>
				{{if .user}}
              		<input class="ipt"   value="{{.user.Name}}" type="text"  name="author" id="author" value="" tabindex="2" placeholder="昵称">
				{{else}}
              		<input class="ipt"  type="text"  name="author" id="author" value="" tabindex="2" placeholder="昵称">
				{{end}}

				<span class="help-inline">昵称 (必填)</span></li>
            <li class="form-inline">
              <label class="hide" for="email">邮箱</label>
              	{{if .user}}
              		<input class="ipt" readonly type="text" name="email" id="email"  value="" tabindex="3" placeholder="{{.user.Email}}">
				{{else}}
              		<input class="ipt" type="text" name="email" id="email" value="" tabindex="3" placeholder="邮箱">
				{{end}}
              
              <span class="help-inline">邮箱 (必填)</span>
          </li>
          </ul>
        </div>
      </div>
    </form>
  </div>
  <div id="postcomments">
    <div id="comments">
      <i class="fa fa-comments-o"></i>
      <b>({{.pagenum}})</b>条回复</div>
    <ol class="commentlist">
		{{range $index, $elem := .comments}}
			<li class="comment odd alt thread-odd thread-alt depth-1" id="comment-14599">
		        <div class="c-avatar">
		          <img alt="" data-original="https://www.gyiscool.com" class="avatar avatar-54 photo" height="54" width="54" src="{{$elem.From_user.HeadImg}}" style="display: block;">
		          <div class="c-main" id="div-comment-14599">{{$elem.Content}}
		            <div class="c-meta">
		              <span class="c-author">
		                <a href="" rel="external nofollow" class="url" target="_blank">{{$elem.From_user.Name}}</a></span>{{$elem.Cdate}}
		              <a rel="nofollow" class="comment-reply-link" href="javascript:;" onclick="return addComment.moveForm( &quot;div-comment-14599&quot;, &quot;{{$elem.Iid}}&quot;, &quot;respond&quot;, &quot;{{$.article.Uid}}&quot; )" aria-label="回复给{{$elem.From_user.Name}}">回复</a></div>
		          </div>
		        </div>
				{{if $elem.To_comment.Content}}
		        <ul class="children">
		          <li class="comment even depth-2" id="comment-18350">
		            <div class="c-avatar">
		              <img alt="" data-original="https://gyiscool.com" srcset="{{$elem.From_user.HeadImg}}" class="avatar avatar-54 photo" height="54" width="54" src="{{$elem.From_user.HeadImg}}" style="display: block;">
		              <div class="c-main" id="div-comment-18350">{{$elem.To_comment.Content}}
		                <div class="c-meta">
		                  <span class="c-author">{{$elem.To_comment.From_user.Name}}</span>{{$elem.To_comment.Cdate}}
		                  <a rel="nofollow" class="comment-reply-link" href="" onclick="return addComment.moveForm( &quot;div-comment-18350&quot;, &quot;{{$elem.To_comment.Iid}}&quot;, &quot;respond&quot;, &quot;{{$.article.Uid}}&quot; )" aria-label="回复给{{$elem.To_comment.From_user.Name}}">回复</a></div>
		              </div>
		            </div>
		          </li>
		        </ul>
				{{end}}
		    </li>
		{{end}}
    </ol>
    <div class="commentnav">
    </div>
  </div>
</div>
</div>
		</div>
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
						<p>这是我的个人博客，如果有问题需要讨论，请评论！</p>
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
			 <a href="" title="gy博客">GY的博客</a> 版权所有，保留一切权利 · <a href="" title="站点地图">站点地图</a>   ·   基于goweb构建   © 2011-2014  ·   托管于 <a rel="nofollow" target="_blank" href="">腾讯云主机</a> &amp; <a rel="nofollow" target="_blank" href="">苏ICP备17004846号-1</a>
			</div>
			<div class="trackcode pull-right"></div>
		</div>
	</footer>


	<div class="rollto">
		<button class="btn btn-inverse" data-type="totop" title="回顶部"><i class="fa fa-arrow-up"></i></button>
		<!--<button class="btn btn-inverse" data-type="torespond" title="发评论"><i class="fa fa-comment-o"></i></button>-->
	</div>
</body>
<div id="cVim-status-bar" style="top: 0px;"></div>
<script type="text/javascript">
		set_pagination_v2({{.pagenum}},{{.page}},10,{{.pageUrl}});
</script>
</html>
