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

type ArtController struct {
	beego.Controller
}

func (c *ArtController) Get() {

	//isAjax := c.Ctx.Input.IsAjax()
	postId := c.Ctx.Input.Param(":id") //获取当前文章的id
	var mystruct *FullRes
	var termss []models.Term

	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("iid", "uid", "title").
		From("term").
		Where("term_id = ?").
		OrderBy("iid").Asc().Limit(20).Offset(0)
	sql := qb.String()

	o.Raw(sql, 0).QueryRows(&termss)

	models.WithProfil(termss, "Uid")

	if postId == "" {

		mystruct = &FullRes{Code: 0, Message: "ok", Count: 0, Terms: termss}
		c.Data["json"] = mystruct

		c.ServeJSON()

	} else {

		post := models.Post{Uid: postId}

		err := o.QueryTable("post").Filter("uid", postId).RelatedSel().One(&post) //

		if err == nil {

			mystruct = &FullRes{Code: 0, Message: "ok", Count: 0, Data: post, Terms: termss}

		} else {

			mystruct = &FullRes{Code: 1, Message: "没有该数据", Count: 0, Terms: termss}

		}

		c.Data["json"] = mystruct
		c.ServeJSON()

	}

}

func (c *ArtController) Delete() {
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

func (c *ArtController) Post() {

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
