package api

import (
	"project/model/errorCode"
	"project/model/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Api struct{}

func (s *Api) Tx(ctx *gin.Context, tx *gorm.DB) {
	if tx.Error != nil {
		response.ErrorWithMsg(tx.Error.Error(), errorCode.OTHER, ctx)
		return
	}
	response.Ok(ctx)
}
