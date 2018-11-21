package main

import (
	"blogByGo/models"
	_ "blogByGo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
)

var globalSessions *session.Manager

//初始化 数据库连接等 session
func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@(127.0.0.1:3306)/gy?charset=utf8", 30)

	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Article))
	orm.RegisterModel(new(models.ArticleContent))
	orm.RegisterModel(new(models.Comment))
	orm.RegisterModel(new(models.Like))
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
		uid := ctx.Input.Session("adminId")
		if uid == nil && ctx.Request.RequestURI != "/adm/login" {
			ctx.Redirect(302, "/adm/login")
		}
	}

	beego.InsertFilter("/adm/*", beego.BeforeRouter, FilterUser)

	beego.Run()
}
