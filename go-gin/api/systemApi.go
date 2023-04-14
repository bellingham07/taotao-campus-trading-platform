package api

import (
	articleApi "com.xpwk/go-gin/api/article"
	commodityApi "com.xpwk/go-gin/api/commodity"
	fileApi "com.xpwk/go-gin/api/file"
	messageApi "com.xpwk/go-gin/api/message"
	orderApi "com.xpwk/go-gin/api/order"
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
	userApi.UserCollectApi
}

type CommodityApi struct {
	commodityApi.CommodityInfoApi
	commodityApi.CommodityHistoryApi
	commodityApi.CommodityCategoryApi
	commodityApi.CommodityCollectApi
}

type OrderApi struct {
	orderApi.OrderCommentApi
	orderApi.OrderInfoApi
}

type FileApi struct {
	fileApi.FileInfoApi
}

type ArticleApi struct {
	articleApi.ArticleCommentApi
	articleApi.ArticleBulletinApi
	articleApi.ArticleContentApi
}

type MessageApi struct {
	messageApi.MessageCommentApi
}
