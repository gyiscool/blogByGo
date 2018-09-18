//总页数 当前页， 每页得条数
function set_pagination(allpage, nowpage, pagesize,url) {
	var str = "<ul>";
	var start = 1;
	var end = allpage;
	if(nowpage > 1){//上一页 省略号 
			//str += "<li class='prev-page'>上一页</li>";
			str += "<li class='prev-page'><a href='"+url+"&page="+(nowpage-1)+"'>上一页</a></li>";
			
	}
	
	if((nowpage-1) >= 3){
		(str += "<li><span> ... </span></li>");
		start = nowpage-2;
	}else if(allpage >=5){//和
		end=5;
	}else if((allpage-nowpage) >= 3){
		end = nowpage+2;
	}
	
	for (var i=start;i<=end;i++)
	{ 
		str += "<li "+((i == nowpage) ? "class='active'": "")+"><a href='"+url+"&page="+(i)+"'>"+i+"</a></li>";
	}
	
	if(nowpage < allpage){//下一页
		((allpage-nowpage) >= 3) && (str += "<li><span> ... </span></li>");
			str += "<li class='next-page'><a href='"+url+"&page="+(nowpage-1)+"'>下一页</a></li>";
	}
	
	str += "</ul>";
	$(".pagination").html(str);

}

//评论分页
function set_pagination_v2(allpage, nowpage, pagesize,url){
	var str = "";
	var start = 1;
	var end = allpage;
	if(nowpage > 1){//上一页 
		str += "<a class='prev page-numbers' href='"+url+"?page="+(nowpage-1)+"#comments'>«</a>";	
	}
	

	if( allpage >= 9){

		if(nowpage >=6){
			str += "<a class='page-numbers' href='"+url+"?page=1#comments'>1</a>"
			str += "<span class='page-numbers'>...</span>"
			start = nowpage-3
		}

		if((end-nowpage) >=5){
			end =nowpage+3
		}
	}
	
	for (var i=start;i<=end;i++)
	{ 
		if(i == nowpage){
			str += "<span class='page-numbers current'>"+i+"</span>";
		}else{
			str += "<a class='page-numbers' href='"+url+"?page="+i+"#comments'>"+i+"</a>";
		}
		
	}

	if( allpage >= 9){
		if((allpage-nowpage) >=5){
			str += "<span class='page-numbers'>...</span>"
			str += "<a class='page-numbers' href='"+url+"?page="+allpage+"#comments'>"+allpage+"</a>"
		}

	}
	
	if(nowpage < allpage){//下一页
		
		str += "<a class='prev page-numbers' href='"+url+"?page="+(nowpage+1)+"#comments'>»</a>";
	}
	

	$(".commentnav").html(str);
	

}
