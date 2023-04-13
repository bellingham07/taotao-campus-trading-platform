package api

import (
	commodityApi "com.xpwk/go-gin/api/commodity"
	"com.xpwk/go-gin/api/user"
)

type SystemApi struct {
	UserApi
	CommodityApi
	FileApi
	OrderApi
	ArticleApi
	MessageApi
}

var SystemApis = new(SystemApi)

type UserApi struct {
	userApi.UserInfoApi
	userApi.UserLocationApi
}

type CommodityApi struct {
	commodityApi.CommodityInfoApi
	commodityApi.CommodityHistoryApi
}
