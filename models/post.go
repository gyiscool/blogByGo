package models

type Post struct {
	Iid      int
	Uid      string `orm:"column(uid);pk"`
	Title    string `orm:"size(100)"`
	Admin    *Admin `orm:"rel(fk);null"`
	Content  string
	Brief    string
	Cdate    string
	Del_flag int
	Utime    string
	Zans     int
	Comments int
	Views    int
	Term     *Term `orm:"null;rel(fk)"`
	Head_img string
}
