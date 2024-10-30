package common

import "project/model"

type Tenant struct {
	model.Model
	Username string `json:"username" form:"username" gorm:"column:username;index"`
	Password string `json:"password" form:"password" gorm:"column:password"`
	NickName string `json:"nickName" form:"nickName" gorm:"column:nick_name"`
}
