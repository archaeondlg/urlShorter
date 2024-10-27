package service

import (
	"project/global"
	"project/model/common"
	"project/utils"
	"time"
)

const (
	size                   = 6         // 短链长度
	expired_duration int64 = 604800000 // 7天 热点数据
)

type ShortUrlService struct {
	Service
}

// 自定义编码
func (s *ShortUrlService) Encode(id int64) string {
	return utils.Int2Base62(id)[:size] // 截断字符串
}

func (s *ShortUrlService) Create(url string, userId int64) error {
	var code string
	id := global.Snowflake.Generate().Int64()
	code = s.Encode(id)
	shortUrl := common.ShortUrl{
		UrlId:     id,
		UserId:    userId,
		Url:       url,
		Code:      code,
		ExpiredAt: time.Now().Unix() + expired_duration,
	}
	tx := global.DB.Create(&shortUrl)
	return tx.Error
}
func (s *ShortUrlService) GetOne(code string) (*common.ShortUrl, error) {
	// 过滤器
	shortUrl := &common.ShortUrl{}
	err := global.DB.Where(&common.ShortUrl{Code: code}).First(shortUrl).Error
	return shortUrl, err
}
