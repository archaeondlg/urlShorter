package service

import (
	"project/global"
	"project/model/common"
	"strings"
	"time"
)

const (
	size = 6
)

var base62List = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
	"l", "m", "n", "o", "p", "q", "r", "s", "t", "u",
	"v", "w", "x", "y", "z", "A",
	"B", "C", "E", "F", "D", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

const bas62Len int64 = 62

type ShortUrlService struct {
	Service
}

// id := global.Snowflake.Generate().Int64()
func (s *ShortUrlService) Encode(id int64) string {
	sb := strings.Builder{}
	for id > 0 {
		index := id % bas62Len
		sb.WriteString(base62List[index])
		id /= bas62Len
	}

	return sb.String()[:size] // 截断字符串
}

var (
	retry = 3
)

const (
	expired_duration int64 = 604800000 // 7天
)

func (s *ShortUrlService) Create(url string, userId int64) error {
	var code string
	id := global.Snowflake.Generate().Int64()
	code = s.Encode(id)
	shortUrl := common.ShortUrl{
		UrlId:     id,
		UserId:    userId,
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
