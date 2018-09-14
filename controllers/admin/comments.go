package admin

import (
	_ "fmt"
	"gojob/models"
	_ "reflect"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CommentsController struct {
	beego.Controller
}

func (c *CommentsController) Get() {

	var pagestart int
	isAjax := c.Ctx.Input.IsAjax()

	if isAjax == true { //ajax 返回
		o := orm.NewOrm()
		page, _ := c.GetInt("page")      //当前页码
		pagesize, _ := c.GetInt("limit") //当前页码
		pagestart = (page - 1) * pagesize

		var posts []models.Comment
		postModel := o.QueryTable("comment")
		nums, _ := postModel.OrderBy("-iid").Count()
		_, _ = postModel.RelatedSel("post", "to_comment", "from_user").OrderBy("-iid").Limit(pagesize, pagestart).All(&posts) //最新修改

		mystruct := &Ress{Code: 0, Message: "记录成功", Count: int(nums), Data: posts}
		c.Data["json"] = mystruct

		c.ServeJSON()

	} else {
		c.Layout = "admin/layout/layout.tpl"
		c.TplName = "admin/comments.tpl"
	}

}
