package admin

import (
	"fmt"
	"blogByGo/models"
	_ "reflect"
	"strconv"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Layout = "admin/layout/layout.tpl"
	c.TplName = "admin/user.tpl"
}

func (c *UserController) Post() {
	var user models.User
	Iid := c.Ctx.Input.Param(":id")

	Id, err := strconv.Atoi(Iid)
	fmt.Print(Id)
	fmt.Print(err)

	o := orm.NewOrm()

	user.Name = c.GetString("Name")
	user.Email = c.GetString("Email")
	user.HeadImg = c.GetString("HeadImg")

	if Iid == "" { //新建

		o.Insert(&user)

	} else {

	}

}

func (c *UserController) Delete() {
	var mystruct *Ress
	var newpost models.User
	postId := c.Ctx.Input.Param(":id") //获取当前文章的id
	o := orm.NewOrm()

	Iid, err := strconv.Atoi(postId)

	if postId == "" || err != nil { //参数有问题
		mystruct = &Ress{Code: 0, Message: "缺少参数或者参数不合法", Count: 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()
	}

	newpost.Id = Iid
	if o.Read(&newpost) == nil {
		newpost.Del_flag = newpost.Id
		_, err := o.Update(&newpost)
		mystruct = &Ress{Code: 0, Message: "更新失败", Count: 0}
		if err == nil {
			mystruct = &Ress{Code: 1, Message: "更新成功", Count: 0}
		}
	} else {

		mystruct = &Ress{Code: 0, Message: "缺少参数", Count: 0}

	}

	c.Data["json"] = mystruct
	c.ServeJSON()

}
