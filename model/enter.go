package model

import (
	"project/global"
	"time"

	"gorm.io/gorm"
)

type CommonModel struct {
	ID        uint           `gorm:"primarykey" json:"Id"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type Model struct{}

func (s *Model) First(query interface{}) (*Model, error) {
	res := global.DB.Where(query).First(s)
	return s, res.Error
}
func (s *Model) Select() {}
func (s *Model) Save() *gorm.DB {
	return global.DB.Save(s)
}
func (s *Model) Update() {}
func (s *Model) Delete() {}
