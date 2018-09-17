package models

type ArticleContent struct {
	Article  *Article `orm:"column(uid);rel(one);pk"`
	Html     string   `orm:"column(html_content)"`
	MarkDown string   `orm:"column(markdown_content)"`
}

func (u *ArticleContent) TableName() string {
	return "article_content"
}
