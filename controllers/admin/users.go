package admin

import (
	"blogByGo/models"
	_ "reflect"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UsersController struct {
	beego.Controller
}

func (c *UsersController) Get() {

	var pagestart int
	isAjax := c.Ctx.Input.IsAjax()
	if isAjax == true { //ajax 返回
		o := orm.NewOrm()
		page, _ := c.GetInt("page")      //当前页码
		pagesize, _ := c.GetInt("limit") //当前页码
		pagestart = (page - 1) * pagesize

		var posts []models.User
		postModel := o.QueryTable("user")
		nums, _ := postModel.Filter("del_flag__iexact", 0).OrderBy("-id").Count()
		_, _ = postModel.Filter("del_flag__iexact", 0).OrderBy("-id").Limit(pagesize, pagestart).All(&posts) //最新修改

		mystruct := &Ress{Code: 0, Message: "记录成功", Count: int(nums), Data: posts}
		c.Data["json"] = mystruct

		c.ServeJSON()

	} else {
		c.Layout = "admin/layout/layout.tpl"
		c.TplName = "admin/users.tpl"
	}

	//返回html

}
