package admin

import (
	"fmt"
	"blogByGo/models"
	_ "reflect"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {

	//isAjax := c.Ctx.Input.IsAjax()
	postId := c.Ctx.Input.Param(":id") //获取当前文章的id
	var termss []models.Term

	o := orm.NewOrm()
	if postId == "" {

		//c.Data[""] = post

		c.Layout = "admin/layout/layout.tpl"
		c.TplName = "admin/article.tpl"
	} else {

		article := models.Article{Uid: postId}

		_ = o.QueryTable("article").Filter("uid", postId).RelatedSel().One(&article) //
		article.Term = &models.Term{Uid: "1"}
		c.Data["article"] = article
		o.LoadRelated(&article, "Content")

		fmt.Printf("参数是%+v\n\n\n", article.Content)

	}

	//获取分类
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("iid", "uid", "title").
		From("term").
		Where("term_id = ?").
		OrderBy("iid").Asc().Limit(20).Offset(0)
	sql := qb.String()

	o.Raw(sql, 0).QueryRows(&termss)

	models.WithProfil(termss, "Uid")

	c.Data["terms"] = termss

	c.Layout = "admin/layout/layout.tpl"
	c.TplName = "admin/article.tpl"

}

func (c *ArticleController) Delete() {
	var mystruct *Ress
	var newArticle models.Article
	postId := c.Ctx.Input.Param(":id") //获取当前文章的id
	o := orm.NewOrm()

	newArticle.Uid = postId
	if postId == "" { //新建
		mystruct = &Ress{Code: 1, Message: "缺少参数", Count: 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()
	}
	if o.Read(&newArticle) == nil {
		newArticle.Del_flag = 1
		_, err := o.Update(&newArticle)
		if err == nil {

			mystruct = &Ress{Code: 1, Message: "删除成功", Count: 0}

		}
	} else {

		mystruct = &Ress{Code: 1, Message: "没有该数据", Count: 0}

	}

	c.Data["json"] = mystruct
	c.ServeJSON()

}

func (c *ArticleController) Post() {

	//isAjax := c.Ctx.Input.IsAjax()
	var post, oldpost models.Article
	var content models.ArticleContent
	var mystruct *Ress

	o := orm.NewOrm()
	adminId := c.GetSession("adminId").(string)
	postId := c.Ctx.Input.Param(":id") //获取当前文章的id
	term := c.GetString("term")

	//article赋值
	post.Title = c.GetString("title")
	post.Term = &models.Term{Uid: term}
	post.Admin = &models.Admin{Uid: adminId}

	//html复制
	content.Html = c.GetString("editormd-html-textarea")
	content.MarkDown = c.GetString("editormd-markdown-textarea")

	if term == "" {
		mystruct = &Ress{Code: 0, Message: "缺少分类", Count: 0}
	}

	post.Term = &models.Term{Uid: term}

	if post.Title == "" {
		mystruct = &Ress{Code: 0, Message: "缺少标题", Count: 0}
	}

	if mystruct != nil {

		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()

	}

	//新增
	if postId == "" {
		timestamp := time.Now().Unix()
		post.Uid = strconv.FormatInt(timestamp, 10)
		post.Cdate = time.Now().Format("2006-01-02 15:04:05")

		//插入主表
		_, postErr := o.Insert(&post)

		if postErr == nil {
			//插入从表
			content.Article = &post
			o.Insert(&content)

			mystruct = &Ress{Code: 0, Message: "添加成功", Data: post, Count: 0}

		} else {

			mystruct = &Ress{Code: 1, Message: "添加失败", Count: 0}
		}

	} else { //更新
		oldpost.Uid = postId
		if o.Read(&oldpost) == nil {
			//post.Uid = postId

			oldpost.Title = c.GetString("title")

			oldpost.Utime = time.Now().Format("2006-01-02 15:04:05")

			_, err := o.Update(&oldpost)

			if err == nil {

				content.Article = &oldpost
				o.Update(&content)
				mystruct = &Ress{Code: 1, Message: "更新成功", Count: 0}
			}
		} else {

			mystruct = &Ress{Code: 1, Message: "没有数据", Count: 0}

		}
	}
	c.Data["json"] = mystruct
	c.ServeJSON()

}
