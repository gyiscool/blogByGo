package models

type User struct {
	Id            int `orm:"pk"`
	Name          string
	Email         string
	HeadImg       string
	From_comments []*Comment `orm:"reverse(many)"`
	Cdate         string
	Del_flag      int
	ZanNum        int
	CommentNum    int
	ViewNum       int
	//To_comments   []*Comment `orm:"reverse(many)"`
}

//多个列表
func WithProfiles(userlist []*User) {
	//var idList  [len(userlist)]int
	idList := make([]int, len(userlist))
	for i, p := range userlist {
		idList[i] = p.Id
	}
}
