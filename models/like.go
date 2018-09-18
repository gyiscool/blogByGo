package models

type Like struct {
	Id       int    `orm:"column(id);pk"`
	UserId   string `orm:"column(user_id)"`
	IsCancle int    `orm:"column(is_cancle)"`
	Article  string `orm:"column(article_id)"`
	Utime    string
	Ctime    string
}
