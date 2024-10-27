package service

import (
	"project/global"
	"project/model/common"
)

type RedirectRecordService struct{}

func (s *RedirectRecordService) Create(record *common.RedirectRecord) error {
	return global.DB.Create(record).Error
}
