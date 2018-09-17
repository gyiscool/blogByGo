package controllers

import (
	"fmt"
	"gojob/models"
	_ "reflect"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {

	var article models.Article
	var preArticle models.Article
	var nextArticle models.Article
	var comments []models.Comment

	page, _ := c.GetInt("page") //当前页码
	pagesize := 1               //每页数量
	pagenum := 0                //最后一页数量
	pagestart := 0              //偏移量

	postId := c.Ctx.Input.Param(":id")

	fmt.Printf("参数是%+v\n\n\n", c.Ctx.Input.Param(":id"))
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

	_ = o.QueryTable("article").Filter("uid", postId).RelatedSel("Admin", "Term").One(&article) //

	fmt.Println("载入")
	o.LoadRelated(&article, "Content")

	_ = o.QueryTable("article").Filter("iid__lt", article.Iid).RelatedSel().OrderBy("-iid").Limit(1).One(&preArticle) //上一个
	_ = o.QueryTable("article").Filter("iid__gt", article.Iid).RelatedSel().OrderBy("iid").Limit(1).One(&nextArticle) //下一个

	var newArticles []models.Article
	num, _ := o.QueryTable("article").RelatedSel().OrderBy("-iid").Limit(6).All(&newArticles)

	//查看评论分页
	commetnModel := o.QueryTable("comment")

	if article.Iid != 0 {
		commetnModel = commetnModel.Filter("post_id", article.Iid)
	}

	nums, _ := commetnModel.RelatedSel().Count()
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

	_, _ = commetnModel.RelatedSel().OrderBy("-iid").Limit(pagesize, pagestart).All(&comments) //最新修改
	fmt.Printf("评论%+v\n\n\n", comments)

	fmt.Printf("查看时间%+v\n\n\n", num)
	fmt.Printf("查看时间%+v\n\n\n", nums)
	c.Data["article"] = article
	c.Data["preArticle"] = preArticle
	c.Data["nextArticle"] = nextArticle
	c.Data["newArticles"] = newArticles
	c.Data["comments"] = comments
	//处理分页
	c.Data["page"] = page                    //当前页码
	c.Data["pagenum"] = pagenum              //总页码
	c.Data["pageUrl"] = "/article/" + postId //当前页码
	c.TplName = "detail.tpl"

}
