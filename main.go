package main

import (
	"fmt"
	"gojob/models"
	_ "gojob/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
)

var globalSessions *session.Manager

//初始化 数据库连接等 session
func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@(127.0.0.1:3306)/gy?charset=utf8", 30)

	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Article))
	orm.RegisterModel(new(models.ArticleContent))
	orm.RegisterModel(new(models.Post))
	orm.RegisterModel(new(models.Profile))
	orm.RegisterModel(new(models.Comment))
	orm.RegisterModel(new(models.Term))
	orm.RegisterModel(new(models.Admin))

	orm.RunSyncdb("default", false, true)

	beego.SetStaticPath("/static", "public")
	beego.SetStaticPath("/images", "images")
	beego.SetStaticPath("/css", "css")
	beego.SetStaticPath("/js", "js")

	beego.BConfig.WebConfig.Session.SessionOn = true

	//globalSessions, _ = session.NewManager("memory",&session.ManagerConfig{CookieName: "gosessionid",EnableSetCookie:true,Gclifetime: 3600,Maxlifetime: 3600,Secure: false,CookieLifeTime: 3600,ProviderConfig: ""})
	globalSessions, _ = session.NewManager("memory", &session.ManagerConfig{CookieName: "gosessionid", Gclifetime: 3600})
	go globalSessions.GC()
}

func main() {

	orm.Debug = true
	//执行 中间件 前置
	var FilterUser = func(ctx *context.Context) {
		//sess, _ := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
		//defer sess.SessionRelease(ctx.ResponseWriter)
		//sess.Set("mySession", "this my session------------")
		uid := ctx.Input.Session("adminId")
		//fmt.Print(se)
		//uid := sess.Get("adminId")
		//fmt.Print("session ")
		//fmt.Println(uid)
		if uid == nil && ctx.Request.RequestURI != "/adm/login" {
			ctx.Redirect(302, "/adm/login")
		}
	}

	//提交评论过滤 匹配email是否合法，
	/*var FilterComment = func(ctx *context.Context) {
		if ctx.BeegoInput.Query("comment") == "" {
			mystruct := &models.Res{Code: 0, Message: "缺少", Data: 0}
			ctx.Data["json"] = mystruct
			ctx.ServeJSON()
		}

		if ctx.BeegoInput.Query("comment_post_ID") == "" {
			mystruct := &models.Res{Code: 0, Message: "缺少", Data: 0}
			ctx.Data["json"] = mystruct
			ctx.ServeJSON()
		}

	}

	beego.InsertFilter("/comment", beego.BeforeRouter, FilterComment)*/
	beego.InsertFilter("/adm/*", beego.BeforeRouter, FilterUser)

	o := orm.NewOrm()
	o.Using("default")
	user := models.User{Id: 1}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name)
		fmt.Printf("年龄是%v\n", user.Profile)

	}
	var userss []*models.User
	o.QueryTable("user").All(&userss)
	// 自动查询到 Profile
	//fmt.Println(userss.Profile)
	//fmt.Println(userss)
	fmt.Printf("所有的列表项是啥： %+v\n", userss)
	models.WithProfiles(userss)
	for _, p := range userss {
		fmt.Printf("列表 %+v\n", p)
		fmt.Printf("每一项 %+v\n", p.Profile)
		//fmt.Printf("每一项 %+v\n", p.Posts)
	}
	// 因为在 Profile 里定义了反向关系的 User，所以 Profile 里的 User 也是自动赋值过的，可以直接取用。

	beego.Run()
}
