package system

import "project/model"

type Admin struct {
	model.Model
	NickName string `json:"nickname" gorm:"column:nickname"`
	Username string `json:"username" gorm:"column:username;"`
	Password string `json:"passwod" gorm:"column:password;"`
	Avatar   string `json:"avatar" gorm:"column:avatar;"`
	Phone    string `json:"phone" gorm:"column:phone;"`
	Email    string `json:"email" gorm:"column:email;"`
	// Enable        int64
	// OriginSetting
	// RoleId   int    `json:"roleId" gorm:"column:role_id"`
}

type AdminLogin struct {
	Username string `json:"username" gorm:"column:username;"`
	Password string `json:"passwod" gorm:"column:password;"`
}

type LoginResponse struct {
	Token     string
	ExpiresAt int64 `json:"expiresAt"`
}
