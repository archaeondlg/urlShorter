package tenant

import (
	"project/api"
	"project/model/common"
	"project/model/errorCode"
	"project/model/response"
	"project/model/system"
	"project/service"
	"project/utils"

	"github.com/gin-gonic/gin"
)

type TenantApi struct {
	api.Api
}

func (s *TenantApi) Login(c *gin.Context) {
	var tenantLogin common.TenantLogin
	c.ShouldBindJSON(&tenantLogin)
	model, err := service.ServiceGroup.TenantService.Login(tenantLogin)
	if err != nil {
		response.ErrorWithMsg(err.Error(), errorCode.OTHER, c)
		return
	}

	baseClaims := utils.BaseClaims{
		ID:     model.ID,
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
