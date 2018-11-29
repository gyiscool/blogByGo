package routers

import (
	"blogByGo/controllers"
	"blogByGo/controllers/admin"
	"blogByGo/controllers/common"

	"github.com/astaxie/beego"
)

func init() {
	//前台
	//展示
	beego.Router("/articles-?:page:int.html", &controllers.MainController{})
	beego.Router("/term/:term_id/articles.html", &controllers.MainController{})
	beego.Router("/term/:term_id/articles-?:page:int.html", &controllers.MainController{})
	beego.Router("/article/?:id.html", &controllers.ArticleController{})

	//前台异步请求
	beego.Router("/comment", &controllers.CommentController{})
	beego.Router("/like/article/:id", &controllers.LikeController{})

	/**后台管理模块
	 */
	beego.Router("/adm/login", &admin.LoginController{})

	beego.Router("/adm/comments", &admin.CommentsController{})
	beego.Router("/adm/comment/:id", &admin.CommentController{})

	beego.Router("/adm/users", &admin.UsersController{})
	beego.Router("/adm/user/?:id", &admin.UserController{})

	beego.Router("/adm/articles", &admin.ArticlesController{})
	beego.Router("/adm/article/?:id", &admin.ArticleController{})

	beego.Router("/adm/art/?:id", &admin.ArtController{})

	beego.Router("/adm/file", &admin.FileController{})
	beego.Router("/common/file", &common.FileController{})

	beego.Router("/api/list", &admin.ArticlesController{})
}
