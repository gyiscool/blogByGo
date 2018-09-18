package controllers

import (
	"blogByGo/models"
	_ "reflect"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LikeController struct {
	beego.Controller
}

func (c *LikeController) Post() {
	var like models.Like
	var article models.Article
	var mystruct *models.SampRes
	var likeNum int64

	uid := c.Ctx.Input.Param(":id") //获取当前文章的id

	userId := c.Ctx.Input.Cookie("beegosessionID")

	o := orm.NewOrm()

	if uid == "" {
		mystruct = &models.SampRes{Success: 0, Message: "没有对应的文章"}
		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()
		//exit
	}

	//是否有文章
	error := o.QueryTable("article").Filter("uid", uid).Filter("del_flag", 0).One(&article) //

	//没有对应文章
	if error == orm.ErrNoRows {
		mystruct = &models.SampRes{Success: 0, Message: "没有对应的文章"}
		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()
		//exit
	}

	//是否赞过
	likeError := o.QueryTable("like").Filter("user_id", userId).Filter("article_id", uid).One(&like) //

	//没有，攒一个，并且统计进去
	if likeError == orm.ErrNoRows {
		like.UserId = userId
		like.IsCancle = 0
		like.Article = uid
		like.Ctime = time.Now().Format("2006-01-02 15:04:05")

		o.Insert(&like)

		//执行统计
		//获取总赞的数量

		likeNum, _ = o.QueryTable("like").Filter("is_cancle", 0).Filter("article_id", uid).Count()

		article.Zans = int(likeNum)
		o.Update(&article)
		mystruct = &models.SampRes{Success: 1, Message: "点赞成功", Data: article.Zans}
		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()

	} else {

		mystruct = &models.SampRes{Success: 0, Message: "赞赞其他文章吧", Data: article.Zans}
		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()

	}

}
