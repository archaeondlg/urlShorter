package model

import (
	"time"

	"gorm.io/gorm"
)

type ModelUpdate struct {
	UpdatedAt time.Time // 更新时间
}

type ModelDelete struct {
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type Model struct {
	ID        uint      `gorm:"primarykey" json:"Id"` // 主键ID
	CreatedAt time.Time // 创建时间
}

type PageQuery struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}
