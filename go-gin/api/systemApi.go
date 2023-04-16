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
	userApi.InfoApi
	userApi.LocationApi
	userApi.CollectApi
}

type CommodityApi struct {
	commodityApi.InfoApi
	commodityApi.HistoryApi
	commodityApi.CategoryApi
	commodityApi.CollectApi
}

type OrderApi struct {
	orderApi.CommentApi
	orderApi.InfoApi
}

type FileApi struct {
	fileApi.InfoApi
}

type ArticleApi struct {
	articleApi.CommentApi
	articleApi.BulletinApi
	articleApi.ContentApi
}

type MessageApi struct {
	messageApi.CommentApi
}
