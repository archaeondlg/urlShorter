package service

import (
	"errors"
	"project/dao"
	"project/global"
	"project/model/system"
	"project/utils"
)

type AdminService struct {
	dao.Dao
}

func (s *AdminService) Login(adminLogin system.AdminLogin) (admin system.Admin, err error) {
	err = global.DB.Where("username = ?", adminLogin.Username).First(&admin).Error
	if err != nil {
		return
	}
	ok := utils.BcryptCheck(adminLogin.Password, admin.Password)
	if !ok {
		return admin, errors.New("账号或密码错误")
	}
	return admin, err
}

func (s *AdminService) ChangePasswd(userId uint, oldPasswd string, newPasswd string) (err error) {
	var user system.Admin
	err = global.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return
	}
	ok := utils.BcryptCheck(oldPasswd, user.Password)
	if !ok {
		return errors.New("原密码错误")
	}
	return global.DB.Model(&user).Update("password", utils.BcryptHash(newPasswd)).Error
}

func (s *AdminService) Profile(userId uint) (admin system.Admin, err error) {
	err = global.DB.Model(&admin).Where("id = ?", userId).First(&admin).Error
	return
}
