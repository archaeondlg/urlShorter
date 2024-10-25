package common

import (
	"project/model"
)

type ShortUrl struct {
	model.Model
	UrlId  int64  `json:"urlId" grom:"column:url_id;primarykey;"`
	UserId int64  `json:"userId" grom:"column:user_id;not null;default:;comment:用户Id;index;"`
	Url    string `json:"url" gorm:"column:url;size:4096;not null;default:;comment:长链接url;"`
	Code   string `json:"code" gorm:"column:code;size:32;not null;default:;comment:短链接url;unique;"`
	// Status    int    `json:"status" gorm:"type:tinyint;not null;default:0;comment:'0:enable;1:disable';"`
	ExpiredAt int64 `json:"ExpiredAt" gorm:"column:expired_at;not null;default:;comment:短链接url;unique;"`
}
