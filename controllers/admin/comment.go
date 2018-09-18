package admin

import (
	_ "fmt"
	"blogByGo/models"
	_ "reflect"
	"strconv"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Delete() {
	var mystruct *Ress
	var newpost models.Comment
	postId := c.Ctx.Input.Param(":id") //获取当前文章的id
	o := orm.NewOrm()

	Iid, err := strconv.Atoi(postId)

	if postId == "" || err != nil { //新建
		mystruct = &Ress{Code: 0, Message: "缺少参数或者参数不合法", Count: 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()
	}

	newpost.Iid = Iid

	if _, err := o.Delete(&newpost); err == nil {

		mystruct = &Ress{Code: 1, Message: "删除成功", Count: 0}
	} else {

		mystruct = &Ress{Code: 0, Message: "没有该数据", Count: 0}
	}

	c.Data["json"] = mystruct
	c.ServeJSON()

}
