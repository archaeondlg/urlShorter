package dao

import (
	"context"

	"gorm.io/gorm"
)

type Cmder interface {
	DB() *gorm.DB
}

type Dao struct {
	Cmder
	Model interface{}
}

// 传入struct或[]struct的指针
func (s *Dao) Create(model interface{}) *gorm.DB {
	return s.DB().Create(model)
}

// value推荐传入map,传入struct会导致零值不更新
func (s *Dao) Update(value interface{}, args ...interface{}) *gorm.DB {
	var query string
	if len(args) > 0 {
		query = args[0].(string)
		args = args[1:]
	}
	return s.DB().Where(query, args...).Updates(value)
}

// 传入where条件
func (s *Dao) Delete(query string, args ...interface{}) *gorm.DB {
	return s.DB().Where(query, args...).Delete(s.Model)
}
func (s *Dao) DeleteByIdList(idList []interface{}) *gorm.DB {
	return s.DB().Where("id in ?", idList).Delete(s.Model)
}

func (s *Dao) GetOne(value interface{}, query string, args ...interface{}) *gorm.DB {
	return s.DB().Where(query, args...).First(value)
}

// 传入where条件, 使用map或字符串
func (s *Dao) GetList(list interface{}, query string, args ...interface{}) *gorm.DB {
	return s.DB().Where(query, args...).Order("id desc").Find(list)
}
func (s Dao) GetPagination(ctx context.Context, list interface{}, query string, args ...interface{}) (total int64, tx *gorm.DB) {
	tx = s.DB().Where(query, args...).Count(&total)
	if tx.Error != nil {
		return
	}
	tx = s.DB().Where(query, args...).Scopes(Paginate(ctx)).Order("id desc").Find(list)
	return
}

func Paginate(ctx context.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := ctx.Value("pageNum").(int)
		if page == 0 {
			page = 1
		}

		pageSize := ctx.Value("pageSize").(int)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
