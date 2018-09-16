package models

type Article struct {
	Iid      int
	Uid      string `orm:"column(uid);pk"`
	Del_flag int
	Admin    *Admin          `orm:"rel(fk);null"`
	Title    string          `orm:"size(100)"`
	Content  *ArticleContent `orm:"reverse(one)"`
	Brief    string
	Cdate    string
	Utime    string
	Zans     int
	Comments int
	Views    int
	Term     *Term `orm:"null;rel(fk)"`
}
