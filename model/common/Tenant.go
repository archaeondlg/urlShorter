package common

import "project/model"

type Tenant struct {
	model.Model
	Username string `json:"username" form:"username" gorm:"column:username;index"`
	Password string `json:"password" form:"password" gorm:"column:password"`
	NickName string `json:"nickName" form:"nickName" gorm:"column:nick_name"`
}

type TenantLogin struct {
	Username string `json:"username" gorm:"column:username;"`
	Password string `json:"password" gorm:"column:password;"`
}

type TenantPassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
