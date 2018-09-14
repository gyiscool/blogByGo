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

	var post models.Post
	var prePost models.Post
	var nextPost models.Post
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

	fmt.Printf("打印主表全部%+v\n\n\n", termss[0])

	models.WithProfil(termss, "Uid")

	c.Data["terms"] = termss

	_ = o.QueryTable("post").Filter("uid", postId).RelatedSel().One(&post) //

	_ = o.QueryTable("post").Filter("iid__lt", post.Iid).RelatedSel().OrderBy("-iid").Limit(1).One(&prePost) //上一个
	_ = o.QueryTable("post").Filter("iid__gt", post.Iid).RelatedSel().OrderBy("iid").Limit(1).One(&nextPost) //下一个
	fmt.Printf("文章的iid%+v\n\n\n", nextPost)
	var newPosts []models.Post
	num, _ := o.QueryTable("post").Exclude("head_img__exact", "").RelatedSel().OrderBy("-iid").Limit(6).All(&newPosts)

	//查看评论分页
	commetnModel := o.QueryTable("comment")

	if post.Iid != 0 {
		commetnModel = commetnModel.Filter("post_id", post.Iid)
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
	c.Data["article"] = post
	c.Data["preArticle"] = prePost
	c.Data["nextArticle"] = nextPost
	c.Data["newArticles"] = newPosts
	c.Data["comments"] = comments
	c.Data["page"] = page                    //当前页码
	c.Data["pagenum"] = pagenum              //总页码
	c.Data["pageUrl"] = "/article/" + postId //当前页码
	c.TplName = "detail.tpl"
}
