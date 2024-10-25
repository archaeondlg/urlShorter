package common

type User struct {
	UserId   int64  `json:"userId" form:"userId" gorm:"column:user_id;primary;"`
	Username string `json:"username" form:"username" gorm:"column:username;index"`
	Password string `json:"password" form:"password" gorm:"column:password"`
	NickName string `json:"nickName" form:"nickName" gorm:"column:nick_name"`
}
