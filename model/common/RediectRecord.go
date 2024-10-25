package common

import "project/model"

type RedirectRecord struct {
	model.Model
	RecordId  int    `json:"record_id" form:"record_id" gorm:"column:record_id;primary;"`
	UrlId     int64  `json:"urlId" gorm:"column:url_id;index;"`
	UserId    int64  `json:"userId" gorm:"column:user_id;not null;default:;comment:用户Id;index;"`
	IP        string `json:"ip" gorm:"column:ip;not null;default:;comment:ip地址;"`
	UA        string `json:"ua" gorm:"column:ua;not null;default:;comment:ua头;"`
	CreatedAt int64  `json:"created_at" form:"created_at" gorm:"column:created_at;"`
}

func (RedirectRecord) TableName() string {
	return "redirect_record"
}
