package router

import (
	"project/api/common"
)

type Group struct {
	PublicRouter
}

var RouterGroup = new(Group)

type ApiGroup struct {
	common.UrlApi
}

var api = new(ApiGroup)
