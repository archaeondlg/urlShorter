package router

import (
	"project/api/common"
	"project/api/system"
	"project/api/tenant"
)

type Group struct {
	PublicRouter
	AdminRouter
	TenantRouter
}

var RouterGroup = new(Group)

type ApiGroup struct {
	common.UrlApi
	system.AdminApi
	tenant.TenantApi
}

var api = new(ApiGroup)
