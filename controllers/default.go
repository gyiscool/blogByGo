package controllers

import (
	"blogByGo/models"
	"blogByGo/models/common"
	"fmt"
	_ "reflect"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	//获取参数
	title := c.GetString("title")
	termId := c.Ctx.Input.Param(":term_id")
	termTitle := ""
	ptermId := termId
	page, _ := c.GetInt("page", 1)

	//定义分页参数 url
	pagesize := 10
	pageUrl := c.Ctx.Input.URL()
	pageOffset := 0
	pagenum := 0

	//其他参数
	var termss []models.Term
	var term models.Term

	var articles []models.Article
	var modelPage *common.ModelPage
	termUrl := "/articles"
	qb, _ := orm.NewQueryBuilder("mysql")

	//构建头部分类标签
	qb.Select("iid", "uid", "title").
		From("term").
		Where("term_id = ?").
		OrderBy("iid").Asc().Limit(20).Offset(0)

	o := orm.NewOrm()
	sql := qb.String()
	o.Raw(sql, 0).QueryRows(&termss)
	models.WithProfil(termss, "Uid")

	//主要列表 以及分页步骤
	articlesModel := o.QueryTable("article")

	//过滤条件 获取分类信息
	if termId != "" {
		articlesModel = articlesModel.Filter("term", termId)

	}

	articlesModel = articlesModel.Filter("is_public", 1)
	articlesModel = articlesModel.Filter("del_flag", 0)

	if title != "" {
		pageUrl = pageUrl + "?title=" + title + "&"
		articlesModel = articlesModel.Filter("title__icontains", title)
	} else {
		pageUrl = pageUrl + "?"
	}

	nums, _ := articlesModel.Count()

	//构建分页条件
	modelPage = &common.ModelPage{TotalNum: int(nums), SearchPage: page, Pagesize: pagesize}
	pageOffset = modelPage.GetOffset()
	pagenum = modelPage.GetTotalPage()

	_, _ = articlesModel.RelatedSel("Admin", "Term").OrderBy("level", "-iid").Limit(pagesize, pageOffset).All(&articles) //最新修改

	var lastArticle []models.Article
	o.QueryTable("article").OrderBy("-iid").Limit(6).All(&lastArticle)

	// 当前分类 信息
	if termId != "" {
		termModel := o.QueryTable("term")
		termModel = termModel.Filter("uid", termId)
		_ = termModel.One(&term) //最新修改

		ptermId = termId
		if term.Term != nil && term.Term.Uid != "" {
			ptermId = term.Term.Uid
		}
		termTitle = term.Title

	}

	c.Data["terms"] = termss
	c.Data["articles"] = articles
	c.Data["newArticles"] = lastArticle
	c.Data["term"] = termId
	c.Data["page"] = page           //当前页码
	c.Data["pagenum"] = pagenum     //总页码
	c.Data["pageUrl"] = pageUrl     //当前页码
	c.Data["termUrl"] = termUrl     //当前页码
	c.Data["ptermId"] = ptermId     //当前页码
	c.Data["termTitle"] = termTitle //当前页码
	fmt.Println(ptermId)

	pagination := &common.Pagination{TotalPage: pagenum, NowPage: page, Pagesize: pagesize, Url: pageUrl}

	fmt.Println(pagination)

	c.Data["string"] = pagination.GetPaginationHtml() //当前页码

	c.Layout = "home/layout/layout.tpl"
	c.TplName = "home/home.tpl"
}
