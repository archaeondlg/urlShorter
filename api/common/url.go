package common

import (
	"net/http"
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
	c.String(http.StatusOK, "Hello %s", code)
	ip := c.ClientIP()
	ua := c.Request.UserAgent()
	shortUrl, err := service.ServiceGroup.ShortUrlService.GetOne(code)
	if err != nil {
		response.ErrorWithMsg("链接不存在", errorCode.OTHER, c)
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
	service.ServiceGroup.RedirectRecordService.Save(&record)
	c.Redirect(errorCode.RedirectTemp, shortUrl.Url)
}

func (s *UrlApi) Create(c *gin.Context) {
	var shortUrl common.ShortUrl
	c.ShouldBindJSON(&shortUrl)
	auth, _ := c.Get("auth")
	authModel := auth.(utils.BaseClaims)
	err := service.ServiceGroup.ShortUrlService.Create(shortUrl.Url, int64(authModel.ID))
	if err != nil {
		response.ErrorWithMsg(err.Error(), errorCode.OTHER, c)
		return
	}
	response.Ok(c)
}
