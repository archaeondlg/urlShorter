package router

import (
	"project/api/common"
	"project/api/system"
)

type Group struct {
	PublicRouter
	AdminRouter
}

var RouterGroup = new(Group)

type ApiGroup struct {
	common.UrlApi
	system.AdminApi
}

var api = new(ApiGroup)
