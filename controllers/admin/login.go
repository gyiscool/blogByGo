package admin

import (
	"fmt"
	"blogByGo/models"
	_ "reflect"
	_ "time"

	"crypto/md5"
	"encoding/hex"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {

	adminId := c.GetSession("adminId")
	fmt.Print(adminId)
	c.TplName = "admin/login.tpl"
}

func (c *LoginController) Post() {

	var mystruct *Ress
	o := orm.NewOrm()
	account := c.GetString("names") //当前页码
	pwd := c.GetString("pwd")       //当前页码

	if account == "" || pwd == "" { //新建
		mystruct = &Ress{Code: 0, Message: "数据不全", Count: 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()
	}

	user := models.Admin{Account: account}
	err := o.Read(&user, "Account")

	if err != nil {
		mystruct = &Ress{Code: 0, Message: "账号或者密码错误", Count: 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()
	}

	fmt.Print(pwd)
	fmt.Print("-----")
	pwd += user.Salt

	h := md5.New()
	h.Write([]byte(pwd))                       // 需要加密的字符串为 sharejs.com
	pwdEncod := hex.EncodeToString(h.Sum(nil)) // 输出加密结果

	fmt.Print(pwd)

	fmt.Print("-----")
	fmt.Print(pwdEncod)

	if pwdEncod != user.Pwd {
		mystruct = &Ress{Code: 0, Message: "账号或者密码错误", Count: 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
		c.StopRun()
	}

	fmt.Print("记录session")
	c.SetSession("adminId", user.Uid)

	c.Data["json"] = &Ress{Code: 1, Message: "登陆成功"}

	c.ServeJSON()
	c.StopRun()

}
