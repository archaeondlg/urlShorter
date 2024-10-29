package system

import "project/model"

type Admin struct {
	model.Model
	NickName string `json:"nickname" gorm:"column:nickname"`
	Username string `json:"username" gorm:"column:username;"`
	Password string `json:"passwod" gorm:"column:password;"`
	Avatar   string `json:"avatar" gorm:"column:avatar;"`
	// RoleId   int    `json:"roleId" gorm:"column:role_id"`
	Phone string `json:"phone" gorm:"column:phone;"`
	Email string `json:"email" gorm:"column:email;"`
	// Enable        int64
	// OriginSetting
}

type AdminLogin struct {
	Username string
	Password string
}

type AdminPassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type LoginResponse struct {
	Token     string
	ExpiresAt int64 `json:"expiresAt"`
}
