package admin

import (
	_ "fmt"
	"gojob/models"
	_ "reflect"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticlesController struct {
	beego.Controller
}

type Ress struct {
	Code    int
	Message string
	Data    interface{}
	Count   int
}

type FullRes struct {
	Code    int
	Message string
	Data    interface{}
	Terms   interface{}
	Count   int
}

type Ress_v2 struct {
	Code    int         `json:"code"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Count   int         `json:"count"`
}

func (c *ArticlesController) Get() {

	var pagestart int
	isAjax := c.Ctx.Input.IsAjax()

	if isAjax == true { //ajax 返回
		o := orm.NewOrm()
		page, _ := c.GetInt("page")      //当前页码
		pagesize, _ := c.GetInt("limit") //当前页码
		pagestart = (page - 1) * pagesize

		var posts []models.Post
		postModel := o.QueryTable("post")
		nums, _ := postModel.Filter("del_flag", 0).RelatedSel().OrderBy("-utime").Count()
		_, _ = postModel.Filter("del_flag", 0).RelatedSel().OrderBy("-utime").Limit(pagesize, pagestart).All(&posts) //最新修改

		mystruct := &Ress{Code: 0, Message: "记录成功", Count: int(nums), Data: posts}
		c.Data["json"] = mystruct

		c.ServeJSON()

	} else {
		c.Layout = "admin/layout/layout.tpl"
		c.TplName = "admin/articles.tpl"

	}

	//返回html

}