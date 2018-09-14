package controllers

import (
	"fmt"
	"gojob/models"
	_ "reflect"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	page, _ := c.GetInt("page")
	term := c.GetString("term")
	title := c.GetString("title")
	pagesize := 10
	pagenum := 0
	pagestart := 0

	termUrl := "/articles"
	pageUrl := "/articles?"

	//构建标签
	var termss []models.Term
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("iid", "uid", "title").
		From("term").
		Where("term_id = ?").
		OrderBy("iid").Asc().Limit(20).Offset(0)
	o := orm.NewOrm()
	sql := qb.String()

	o.Raw(sql, 0).QueryRows(&termss)

	models.WithProfil(termss, "Uid")

	c.Data["terms"] = termss

	var posts []models.Post
	postModel := o.QueryTable("post")

	if term != "" {
		pageUrl = pageUrl + "&term=" + term
		postModel = postModel.Filter("term", term)
	}

	if title != "" {
		pageUrl = pageUrl + "&title=" + title
		postModel = postModel.Filter("title__icontains", title)
	}

	nums, _ := postModel.RelatedSel().OrderBy("-utime").Count()
	fmt.Printf("最新%+v\n\n\n", page)

	if page != 0 || page > 0 {
		pagestart = (page - 1) * pagesize
	} else {
		page = 1
	}

	//pageUrl = pageUrl + "&page=" + string(page)

	fmt.Printf("当前页码%+v\n\n\n", page)
	pagenum = int(int(nums) / int(pagesize))

	if int(pagesize*pagenum) > int(nums) {
		pagenum++
	}

	fmt.Printf("开始页数%+v\n\n\n", pagestart)
	fmt.Printf("页码总数%+v\n\n\n", pagenum)
	fmt.Printf("页码%+v\n\n\n", pagesize)
	_, _ = postModel.RelatedSel().OrderBy("-utime").Limit(pagesize, pagestart).All(&posts) //最新修改

	var newPosts []models.Post
	num, _ := o.QueryTable("post").Exclude("head_img__exact", "").RelatedSel().OrderBy("-iid").Limit(6).All(&newPosts)
	fmt.Printf("开始页数%+v\n\n\n", num)
	c.Data["articles"] = posts
	c.Data["newArticles"] = newPosts
	c.Data["term"] = term
	c.Data["page"] = page       //当前页码
	c.Data["pagenum"] = pagenum //总页码
	c.Data["pageUrl"] = pageUrl //当前页码
	c.Data["termUrl"] = termUrl //当前页码

	c.TplName = "home.tpl"
}

/*t, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
fmt.Printf("查看时间%+v\n\n\n", t)
fmt.Printf("查看时间%+v\n\n\n", t.Year()) //t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
*/
