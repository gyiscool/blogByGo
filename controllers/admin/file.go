package admin

import (
	"fmt"
	"gojob/models"
	_ "reflect"
	_ "time"

	"math/rand"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) Get() {

	isAjax := c.Ctx.Input.IsAjax()

	if isAjax == true { //ajax 返回

		var res [1]models.Returns
		res[0] = models.Returns{City: "北京", Classify: "1", Experience: "1", Id: "2", Logins: "1", Score: "1", Sex: "1", Sign: "1", Username: "1", Wealth: "1"}
		mystruct := &Ress{Code: 0, Message: "记录成功", Count: 100, Data: res}
		c.Data["json"] = mystruct

		c.ServeJSON()

	} else {
		c.Layout = "admin/layout/layout.tpl"
		c.TplName = "admin/articles.tpl"

	}

	//返回html

}

func (c *FileController) Post() {

	f, h, err := c.GetFile("file")
	r := rand.Intn(10)

	fmt.Println("getfile err ", r)
	timestamp := time.Now().Unix()
	var fileSrc string
	fileNameNot := strconv.FormatInt(timestamp, 10)
	fileSrc = `public\file\`
	//var fileName = (string(fileSrc) + string(r) + strconv.FormatInt(timestamp, 10) + string(h.Filename))
	var fileName1 = (string(fileSrc) + strconv.FormatInt(timestamp, 10) + string(path.Ext(h.Filename)))

	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		c.SaveToFile("file", fileName1)
	}

	var res models.FileRes

	res.FileUrl = ("/static/file/" + fileNameNot + string(path.Ext(h.Filename)))
	res.FileSrc = fileName1
	res.FielName = (fileNameNot + string(path.Ext(h.Filename)))
	res.Src = (fileNameNot + string(path.Ext(h.Filename)))
	res.Size = "12"
	mystruct := &Ress_v2{Code: 0, Message: "记录成功", Count: 100, Data: res}
	c.Data["json"] = mystruct

	c.ServeJSON()
}
