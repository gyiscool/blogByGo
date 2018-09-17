package common

import (
	"fmt"
	"strconv"
)

type Pagination struct {
	TotalPage int
	NowPage   int
	Pagesize  int
	Url       string
}

func (pagination *Pagination) GetPaginationHtml() string {

	//共十一个 元素
	var str = ""
	start := 1
	end := pagination.TotalPage
	url := pagination.Url
	nowpage := pagination.NowPage
	prePage := strconv.Itoa(pagination.NowPage - 1)
	nextPage := strconv.Itoa(pagination.NowPage + 1)

	if pagination.TotalPage <= 1 {
		return str
	}
	//上一页
	if pagination.NowPage > 1 {
		str += "<li class='prev-page'><a href='" + url + "&page=" + (prePage) + "'>上一页</a></li>"
	}

	if end >= 9 {

		fmt.Println("9")
		//前省略号 以及第一页
		if nowpage >= 6 {

			fmt.Println("7")

			str += "<li><a href='" + pagination.Url + "?page=1'>1</a> </li>"

			str += "<li><span> ... </span> </li>"
			start = nowpage - 3
			fmt.Println(str)

		}

		//后省略号
		if (pagination.TotalPage - pagination.NowPage) >= 5 {
			end = nowpage + 3
		}

		//是否有下一页
		if pagination.TotalPage != nowpage {
		}

	}

	for i := start; i <= end; i++ {
		if i == nowpage {
			str += "<li class='active'><a href='" + url + "&page=" + strconv.Itoa(i) + "'>" + strconv.Itoa(i) + "</a></li>"
		} else {
			str += "<li><a href='" + url + "&page=" + strconv.Itoa(i) + "'>" + strconv.Itoa(i) + "</a></li>"
		}
	}

	if pagination.TotalPage >= 9 {

		//后省略号 以及最后一页
		if (pagination.TotalPage - pagination.NowPage) >= 5 {

			str += "<li><span>...</span></li>"

			str += "<li><a href='" + url + "&page=" + strconv.Itoa(pagination.TotalPage) + "'>" + strconv.Itoa(pagination.TotalPage) + "</a></li>"

			end = nowpage + 5

		}

	}

	if nowpage < pagination.TotalPage { //下一页
		str += "<li class='next-page'><a href='" + url + "&page=" + (nextPage) + "'>下一页</a></li>"
	}

	return str
}
