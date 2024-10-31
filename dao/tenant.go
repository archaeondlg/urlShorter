package dao

import (
	"project/global"
	"project/model/common"

	"gorm.io/gorm"
)

type TenantDao struct {
	Dao
}

func (s *TenantDao) DB() *gorm.DB {
	model := common.Tenant{}
	return global.DB.Model(model)
}
