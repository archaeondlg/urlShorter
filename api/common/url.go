package common

import (
	"project/model/common"
	"project/model/errorCode"
	"project/model/response"
	"project/service"
	"project/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type UrlApi struct{}

func (s *UrlApi) Access(c *gin.Context) {
	code := c.Param("code")
	ip := c.ClientIP()
	ua := c.Request.UserAgent()
	shortUrl, err := service.ServiceGroup.ShortUrlService.GetOne(code)
	if err != nil {
		response.ErrorWithMsg("链接不存在", errorCode.OTHER, c)
		return
	}

	now := time.Now().Unix()
	if now >= shortUrl.ExpiredAt {
		response.ErrorWithMsg("链接已失效", errorCode.Invalid, c)
		return
	}
	record := common.RedirectRecord{
		UrlId:  shortUrl.UrlId,
		UserId: shortUrl.UserId,
		IP:     ip,
		UA:     ua,
	}
	err = service.ServiceGroup.RedirectRecordService.Create(&record)
	if err != nil {
		response.ErrorWithMsg(err.Error(), errorCode.OTHER, c)
		return
	}

	c.Redirect(errorCode.RedirectTemp, shortUrl.Url)
}

func (s *UrlApi) Create(c *gin.Context) {
	var shortUrl common.ShortUrl
	c.ShouldBindJSON(&shortUrl)
	authModel := utils.Auth(c)
	shortUrl, err := service.ServiceGroup.ShortUrlService.Create(shortUrl.Url, int64(authModel.ID))
	if err != nil {
		response.ErrorWithMsg(err.Error(), errorCode.OTHER, c)
		return
	}
	response.OkWithData(shortUrl, c)
}
