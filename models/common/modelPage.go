package common

type ModelPage struct {
	TotalNum   int
	SearchPage int
	Pagesize   int
}

func (page *ModelPage) GetTotalPage() int {

	nums := page.TotalNum
	pagesize := page.Pagesize

	//整除取最后一个完整pageSize的页码
	pagenum := int(int(nums) / int(pagesize))

	//如果需要加残余元素的页码
	if int(pagesize*pagenum) > int(nums) {
		pagenum++
	}

	return pagenum
}

//获取偏移量
func (page *ModelPage) GetOffset() int {

	pagestart := 0
	pagesize := page.Pagesize

	if page.SearchPage >= 0 {

		pagestart = (page.SearchPage - 1) * pagesize

	}

	return pagestart
}

func (page *ModelPage) GetLimit() int {
	return page.Pagesize
}
