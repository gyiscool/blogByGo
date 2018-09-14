package admin

import (
	"fmt"
	"gojob/models"
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

		post := models.Post{Uid: postId}

		_ = o.QueryTable("post").Filter("uid", postId).RelatedSel().One(&post) //

		fmt.Print(post.Term.Uid)
		fmt.Print(post.Term)
		c.Data["article"] = post

		fmt.Printf("参数是%+v\n\n\n", postId)

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

	fmt.Printf("打印主表全部%+v\n\n\n", termss[0])

	models.WithProfil(termss, "Uid")

	c.Data["terms"] = termss

	c.Layout = "admin/layout/layout.tpl"
	c.TplName = "admin/article.tpl"

}

func (c *ArticleController) Delete() {
	var mystruct *Ress
	var newpost models.Post
	postId := c.Ctx.Input.Param(":id") //获取当前文章的id
	o := orm.NewOrm()

	newpost.Uid = postId
	if postId == "" { //新建
		mystruct = &Ress{Code: 1, Message: "缺少参数", Count: 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()
	}
	if o.Read(&newpost) == nil {
		newpost.Del_flag = 1
		_, err := o.Update(&newpost)
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
	var post, oldpost models.Post
	var mystruct *Ress
	o := orm.NewOrm()
	fmt.Println("开始")

	adminId := c.GetSession("adminId").(string)
	postId := c.Ctx.Input.Param(":id") //获取当前文章的id

	fmt.Println(postId)
	term_id := c.GetString("Term_id")

	post.Title = c.GetString("Title")
	post.Content = c.GetString("Content")
	post.Brief = c.GetString("Brief")
	post.Content = c.GetString("Content")
	post.Head_img = c.GetString("Headimg")

	post.Term = &models.Term{Uid: term_id}
	post.Admin = &models.Admin{Uid: adminId}

	if term_id == "" {
		mystruct = &Ress{Code: 0, Message: "缺少分类", Count: 0}
	}

	post.Term = &models.Term{Uid: term_id}

	if post.Title == "" {
		mystruct = &Ress{Code: 0, Message: "缺少题目", Count: 0}
	}
	if post.Content == "" {
		mystruct = &Ress{Code: 0, Message: "缺少内容", Count: 0}
	}
	if post.Brief == "" {
		mystruct = &Ress{Code: 0, Message: "缺少简介", Count: 0}
	}

	if post.Head_img == "" {
		//	mystruct = &Ress{Code: 0, Message: "缺少概略图", Count: 0}
	}

	fmt.Println("判断")
	fmt.Println(post)
	if mystruct != nil {

		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()
	}

	if postId == "" { //新建
		timestamp := time.Now().Unix()
		post.Uid = strconv.FormatInt(timestamp, 10)
		post.Cdate = time.Now().Format("2006-01-02 15:04:05")
		o.Insert(&post)

		_, postErr := o.Insert(&post)

		if postErr == nil {
			mystruct = &Ress{Code: 0, Message: "添加失败", Count: 0}
		} else {
			mystruct = &Ress{Code: 1, Message: "添加成功", Count: 0}
		}

	} else { //更新
		oldpost.Uid = postId
		if o.Read(&oldpost) == nil {
			//post.Uid = postId

			oldpost.Title = c.GetString("Title")
			oldpost.Content = c.GetString("Content")
			oldpost.Brief = c.GetString("Brief")
			oldpost.Content = c.GetString("Content")
			oldpost.Head_img = c.GetString("Headimg")
			oldpost.Utime = time.Now().Format("2006-01-02 15:04:05")

			_, err := o.Update(&oldpost)

			if err == nil {
				mystruct = &Ress{Code: 1, Message: "更新成功", Count: 0}
			}
		} else {

			mystruct = &Ress{Code: 1, Message: "没有该数据", Count: 0}

		}
	}
	c.Data["json"] = mystruct
	c.ServeJSON()

}
