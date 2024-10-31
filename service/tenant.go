package service

import (
	"errors"
	"project/dao"
	"project/model"
	"project/model/common"
	"project/utils"

	"gorm.io/gorm"
)

type TenantService struct {
	dao.TenantDao
}

func (s *TenantService) Register() {

}

func (s *TenantService) Login(form common.TenantLogin) (model common.Tenant, err error) {
	tx := s.GetOne(&model, "username = ?", form.Username)
	if tx.Error != nil {
		return model, tx.Error
	}
	ok := utils.BcryptCheck(form.Password, model.Password)
	if !ok {
		return model, errors.New("账号或密码错误")
	}
	return
}

func (s *TenantService) Create(model *common.Tenant) *gorm.DB {
	return s.Dao.Create(model)
}

// 修改密码
func (s *TenantService) ChangePasswd(modelId string, change model.ChangePassword) (err error) {
	var model common.Tenant
	tx := s.GetOne(&model, "id = ?", modelId)
	if tx.Error != nil {
		return
	}
	ok := utils.BcryptCheck(model.Password, change.OldPassword)
	if !ok {
		return errors.New("旧密码错误")
	}
	model.Password = utils.BcryptHash(change.NewPassword)
	return s.Update(&model).Error
}
