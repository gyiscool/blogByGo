package controllers

import (
	"fmt"
	"gojob/models"
	_ "gojob/models"
	_ "reflect"
	_ "time"

	"regexp"
	"time"

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
	var newUser models.User

	var comment models.Comment
	var post models.Post
	o := orm.NewOrm()
	content := c.GetString("comment")        //获取内容
	postId := c.GetString("comment_post_ID") //当前页码
	//parentId := c.GetString("comment_parent") //获取的id
	author := c.GetString("author")                  //获取的id
	email := c.GetString("email")                    //获取的id
	toId, commentError := c.GetInt("comment_parent") //获取的id

	//查找是否有 对应的文章
	_ = o.QueryTable("post").Filter("uid", postId).RelatedSel().One(&post) //

	if content == "" {
		mystruct := &models.Res{Code: 0, Message: "缺少内容", Data: 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
	}

	if postId == "" {
		mystruct := &models.Res{Code: 0, Message: "缺少目标", Data: 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
	}

	//检查是否有session
	v := c.GetSession("uid")

	//写入存储
	/**
	  1-用户登陆状态记录
	*/
	if v == nil {
		//没有email 和 author 报错
		if email == "" {
			mystruct := &models.Res{Code: 0, Message: "缺少邮箱", Data: 0}
			c.Data["json"] = mystruct
			c.ServeJSON()
		}

		reg := regexp.MustCompile("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$")
		if !reg.MatchString(email) {
			mystruct := &models.Res{Code: 0, Message: "邮箱不合法", Data: 0}
			c.Data["json"] = mystruct
			c.ServeJSON()
		}

		//是否有过这个人 有写入session 没有增加记录

		_ = o.QueryTable("user").Filter("email", email).One(&user)

		fmt.Printf("准备插入数据%+v\n\n\n", c.Data["Email"])
		fmt.Printf("准备插入数据%+v\n\n\n", user)
		if user.Id == 0 { //新增

			newUser.Name = author
			newUser.Email = email
			newUser.Cdate = time.Now().Format("2006-01-02 15:04:05")
			newUser.Profile = &models.Profile{Id: 0}

			Uid, _ := o.Insert(&newUser)

			c.SetSession("uid", newUser.Id)

			fmt.Printf("写入数据session%+v\n", newUser)
			fmt.Printf("写入数据session%+v\n", Uid)
		} else { //有数据 记录数据
			if author != "" && user.Name != author { //更改数据
				user.Name = author
				o.Update(&user)
			}

			c.SetSession("uid", user.Id)
			fmt.Printf("产看123session%+v\n", user.Id)
		}

	} else { //存在session

		user.Id = v.(int)

		_ = o.Read(&user)

		fmt.Printf("参数是session%+v\n\n\n", user)
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

	mystruct := &models.Res{Code: 0, Message: "记录成功", Data: nil}
	c.Data["json"] = mystruct
	fmt.Printf("参数是%+v\n\n\n", c.Data["Email"])

	c.ServeJSON()
}
