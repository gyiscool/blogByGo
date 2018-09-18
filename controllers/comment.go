package controllers

import (
	"fmt"
	"blogByGo/models"
	_ "blogByGo/models"
	_ "reflect"
	"strconv"
	_ "time"

	"regexp"
	"time"

	"math/rand"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CommentController struct {
	beego.Controller
}

type commetForm struct {
	Nick     string `form:"username"`
	Email    string
	Comment  string
	postId   string
	parentId string
}

func (c *CommentController) Post() {

	var user models.User

	var comment models.Comment
	var post models.Article
	o := orm.NewOrm()
	content := c.GetString("comment")        //获取内容
	postId := c.GetString("comment_post_ID") //当前页码
	//parentId := c.GetString("comment_parent") //获取的id
	author := c.GetString("author")                  //获取的id
	email := c.GetString("email")                    //获取的id
	toId, commentError := c.GetInt("comment_parent") //获取的id

	//查找是否有 对应的文章
	_ = o.QueryTable("article").Filter("uid", postId).RelatedSel().One(&post) //

	if content == "" {
		mystruct := &models.SampRes{Success: 0, Message: "还没有填写回复内容呢^_^", Data: 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
	}

	if postId == "" {
		mystruct := &models.SampRes{Success: 0, Message: "文章参数错误", Data: 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
	}

	//检查是否有session
	v := c.GetSession("uid")

	/**
	  1-用户登陆状态记录
	*/
	if v == nil {
		//没有email 和 author 报错
		if email == "" {
			mystruct := &models.SampRes{Success: 0, Message: "你忘记填邮箱了吧^.^", Data: 0}
			c.Data["json"] = mystruct
			c.ServeJSON()
			c.StopRun()
		}

		if author == "" {
			mystruct := &models.SampRes{Success: 0, Message: "你忘记填昵称了吧^.^", Data: 0}
			c.Data["json"] = mystruct
			c.ServeJSON()
			c.StopRun()
		}

		reg := regexp.MustCompile("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$")
		if !reg.MatchString(email) {
			mystruct := &models.SampRes{Success: 0, Message: "再检查一遍是否是合法邮箱哦@……@", Data: 0}
			c.Data["json"] = mystruct
			c.ServeJSON()
			c.StopRun()
		}

		//是否有过这个人 有写入session 没有增加记录
		error := o.QueryTable("user").Filter("email", email).One(&user)

		if error == orm.ErrNoRows { //新增

			user.Name = author
			user.Email = email
			user.Cdate = time.Now().Format("2006-01-02 15:04:05")
			user.HeadImg = "/static/img/headimg/" + strconv.Itoa(rand.Intn(15)) + ".jpg"

			id, _ := o.Insert(&user)
			fmt.Println("查看")
			fmt.Println(user)
			user.Id = int(id)

			c.SetSession("uid", user.Id)

		} else { //有数据 记录数据
			if author != "" && user.Name != author { //更改数据
				user.Name = author
				o.Update(&user)
			}

			c.SetSession("uid", user.Id)
		}

	} else { //存在session

		user.Id = v.(int)

		o.Read(&user)

	}

	/**
	  2-记录回复内容
	*/
	comment.Cdate = time.Now().Format("2006-01-02 15:04:05")
	comment.Post = &post
	comment.Content = content
	if toId != 0 && commentError == nil {
		comment.To_comment = &models.Comment{Iid: toId}
	} else {
		comment.To_comment = &models.Comment{Iid: 0}
	}
	comment.From_user = &user
	o.Insert(&comment)

	mystruct := &models.SampRes{Success: 1, Message: "记录成功", Data: user}
	c.Data["json"] = mystruct

	c.ServeJSON()
}
