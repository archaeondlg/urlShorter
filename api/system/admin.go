package system

import (
	"project/model/errorCode"
	"project/model/response"
	"project/model/system"
	"project/service"
	"project/utils"

	"github.com/gin-gonic/gin"
)

type AdminApi struct{}

func (s *AdminApi) Login(c *gin.Context) {
	var adminLogin system.AdminLogin
	c.ShouldBindJSON(&adminLogin)
	admin, err := service.ServiceGroup.AdminService.Login(adminLogin)
	if err != nil {
		response.ErrorWithMsg(err.Error(), errorCode.OTHER, c)
		return
	}
	baseClaims := utils.BaseClaims{
		ID:     admin.ID,
		RoleId: 1,
	}
	token, err := utils.NewToken(baseClaims)
	if err != nil {
		response.ErrorWithMsg(err.Error(), errorCode.OTHER, c)
		return
	}
	response.OkWithData(system.LoginResponse{
		Token: token,
	}, c)
}

func (s *AdminApi) ChangePasswd(c *gin.Context) {
	var change system.AdminPassword
	c.ShouldBindJSON(&change)
	admin := utils.Auth(c)
	err := service.ServiceGroup.AdminService.ChangePasswd(admin.ID, change.OldPassword, change.NewPassword)
	if err != nil {
		response.ErrorWithMsg(err.Error(), errorCode.OTHER, c)
		return
	}
	response.Ok(c)
}

}
