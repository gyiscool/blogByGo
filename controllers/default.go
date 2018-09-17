package controllers

import (
	"gojob/models"
	"gojob/models/common"
	_ "reflect"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	termId := c.Ctx.Input.Param(":term_id")

	page, _ := c.GetInt("page")
	title := c.GetString("title")
	pagesize := 1
	pagenum := 0
	pagestart := 0

	termUrl := "/articles"
	pageUrl := "?"

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

	var articles []models.Article
	articlesModel := o.QueryTable("article")

	if termId != "" {
		pageUrl = pageUrl + "&term=" + termId
		articlesModel = articlesModel.Filter("term", termId)
	}

	if title != "" {
		pageUrl = pageUrl + "&title=" + title
		articlesModel = articlesModel.Filter("title__icontains", title)
	}

	nums, _ := articlesModel.OrderBy("-utime").Count()

	if page != 0 || page > 0 {
		pagestart = (page - 1) * pagesize
	} else {
		page = 1
	}

	pagenum = int(int(nums) / int(pagesize))

	if int(pagesize*pagenum) > int(nums) {
		pagenum++
	}

	_, _ = articlesModel.RelatedSel("Admin", "Term").OrderBy("-utime").Limit(pagesize, pagestart).All(&articles) //最新修改

	var lastArticle []models.Article
	o.QueryTable("article").OrderBy("-iid").Limit(6).All(&lastArticle)

	c.Data["articles"] = articles
	c.Data["newArticles"] = lastArticle
	c.Data["term"] = termId
	c.Data["page"] = page       //当前页码
	c.Data["pagenum"] = pagenum //总页码
	c.Data["pageUrl"] = pageUrl //当前页码
	c.Data["termUrl"] = termUrl //当前页码

	pagination := &common.Pagination{TotalPage: pagenum, NowPage: page, Pagesize: pagesize, Url: "?"}

	c.Data["string"] = pagination.GetPaginationHtml() //当前页码
	//fmt.Println(c.Data["string"])

	c.TplName = "home.tpl"
}
