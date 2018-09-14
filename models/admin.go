package models

type Admin struct {
	Iid       int
	Uid       string `orm:"column(uid);pk"`
	Account   string `orm:"size(100)"`
	Nick_name string
	Pwd       string
	Salt      string
	Cdate     string
	Del_flag  int
	Update    string
}
