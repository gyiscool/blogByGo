package models

import (
	"fmt"
)

type User struct {
	Id            int `orm:"pk"`
	Name          string
	Email         string
	HeadImg       string
	Profile       *Profile   `orm:"null;rel(one)"` // OneToOne relation
	From_comments []*Comment `orm:"reverse(many)"`
	Cdate         string
	Del_flag      int
	ZanNum        int
	CommentNum    int
	ViewNum       int
	//To_comments   []*Comment `orm:"reverse(many)"`
}

//利用反射试试看 单个的
func (user *User) WithProfile() {
	fmt.Printf("输出看看%+v", user)
}

//多个列表
func WithProfiles(userlist []*User) {
	//var idList  [len(userlist)]int
	idList := make([]int, len(userlist))
	for i, p := range userlist {
		idList[i] = p.Id
		fmt.Printf("用户id:%+v", i)
		fmt.Printf("用户详情id:%+v", p)
	}
}
