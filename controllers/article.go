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

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {

	var article models.Article
	var preArticle models.Article
	var nextArticle models.Article
	var comments []models.Comment
	var modelPage *common.ModelPage
	var pageOffset int
	var user models.User

	page, _ := c.GetInt("page") //当前页码
	pagesize := 5               //每页数量
	pagenum := 0                //最后一页数量

	postId := c.Ctx.Input.Param(":id")

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
	article.Views += 1
	o.Update(&article)

	o.LoadRelated(&article, "Content")

	//上一个，下一个
	_ = o.QueryTable("article").Filter("iid__lt", article.Iid).RelatedSel().OrderBy("-iid").Limit(1).One(&preArticle) //上一个
	_ = o.QueryTable("article").Filter("iid__gt", article.Iid).RelatedSel().OrderBy("iid").Limit(1).One(&nextArticle) //下一个

	//右侧
	var newArticles []models.Article
	o.QueryTable("article").RelatedSel().OrderBy("-iid").Limit(6).All(&newArticles)

	//查看评论分页
	commetnModel := o.QueryTable("comment")

	if article.Iid != 0 {
		commetnModel = commetnModel.Filter("post_id", article.Uid)
	}

	nums, _ := commetnModel.RelatedSel().Count()

	modelPage = &common.ModelPage{TotalNum: int(nums), SearchPage: page, Pagesize: pagesize}
	pageOffset = modelPage.GetOffset()
	pagenum = modelPage.GetTotalPage()

	_, _ = commetnModel.RelatedSel().OrderBy("-iid").Limit(pagesize, pageOffset).All(&comments) //最新修改

	//查找个人信息
	//检查是否有session
	v := c.GetSession("uid")
	if v != nil {
		user.Id = v.(int)

		o.Read(&user)

		c.Data["user"] = user

	}

	fmt.Println("是个啥")
	fmt.Println(user)
	c.Data["article"] = article
	c.Data["preArticle"] = preArticle
	c.Data["nextArticle"] = nextArticle
	c.Data["newArticles"] = newArticles
	c.Data["comments"] = comments
	//处理分页
	c.Data["page"] = page                    //当前页码
	c.Data["pagenum"] = pagenum              //总页码
	c.Data["pageUrl"] = "/article/" + postId //当前页码

	c.Layout = "home/layout/layout.tpl"
	c.TplName = "home/detail.tpl"

}
