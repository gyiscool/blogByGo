package common

import (
	"fmt"
	_ "gojob/models"
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

type Ress_v2 struct {
	Success    int         `json:"success"`
	Message string      `json:"message"`
	Url    interface{} `json:"url"`
	
}

//对IP做限制
func (c *FileController) Post() {

	f, h, err := c.GetFile("editormd-image-file")
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
		c.SaveToFile("editormd-image-file", fileName1)
	}

	mystruct := &Ress_v2{Success: 1, Message: "记录成功", Url: ("/static/file/" + fileNameNot + string(path.Ext(h.Filename)))}
	c.Data["json"] = mystruct

	c.ServeJSON()
}
