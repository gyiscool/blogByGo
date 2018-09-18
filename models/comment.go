package models

type Comment struct {
	Iid        int      `orm:"column(Iid);pk"`
	Post       *Article `orm:"rel(fk)"`
	Content    string
	To_comment *Comment   `orm:"column(to_uid);null;rel(fk)"`
	From_user  *User      `orm:"column(from_uid);null;rel(fk)"`
	Comments   []*Comment `orm:"reverse(many)"`
	Cdate      string
}
