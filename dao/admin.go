package dao

import (
	"project/global"
	"project/model/system"

	"gorm.io/gorm"
)

type AdminDao struct {
	Dao
}

func (s *AdminDao) DB() *gorm.DB {
	model := system.Admin{}
	return global.DB.Model(model)
}
