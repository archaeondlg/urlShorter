package response

import (
	"net/http"
	"project/model/errorCode"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Json(code int64, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Ok(c *gin.Context) {
	Json(errorCode.OK, map[string]interface{}{}, "", c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Json(errorCode.OK, data, "", c)
}

func OkWithMsg(msg string, data interface{}, c *gin.Context) {
	Json(errorCode.OK, data, msg, c)
}

func Error(code int64, c *gin.Context) {
	Json(code, map[string]interface{}{}, "error", c)
}

func ErrorWithMsg(msg string, code int64, c *gin.Context) {
	Json(code, map[string]interface{}{}, msg, c)
}

func ErrorWithData(data interface{}, msg string, code int64, c *gin.Context) {
	Json(code, data, msg, c)
}
