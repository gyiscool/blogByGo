package models

import (
	_ "fmt"

	"github.com/astaxie/beego/orm"
)

type Term struct {
	Iid   int
	Uid   string `orm:"column(uid);pk"`
	Title string
	Tree  int
	Cdate string
	Pid   string
	Terms []*Term    `orm:"reverse(many)"`
	Term  *Term      `orm:"rel(fk)"` //关联个毛线
	Posts []*Article `orm:"reverse(many)"`
}

//list的关联查询
func WithProfil(mainTermList []Term, fieldStr string) {

	idList := ArrayColums(mainTermList, fieldStr)
	o := orm.NewOrm()

	var userss []Term
	o.QueryTable("term").Filter("term_id__in", idList).All(&userss)

	for i, main := range mainTermList {
		//j := 0
		for j, m := range userss {
			if m.Term.Uid == main.Uid {
				mainTermList[i].Terms = append(mainTermList[i].Terms, &userss[j])
			}
		}
	}

}
