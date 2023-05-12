package api

import (
	articleApi "com.xpdj/go-gin/api/article"
	commodityApi "com.xpdj/go-gin/api/commodity"
	fileApi "com.xpdj/go-gin/api/file"
	messageApi "com.xpdj/go-gin/api/message"
	orderApi "com.xpdj/go-gin/api/order"
	userApi "com.xpdj/go-gin/api/user"
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
	commodityApi.TagApi
	commodityApi.CollectApi
}

type OrderApi struct {
	orderApi.CommentApi
	orderApi.InfoApi
}

type FileApi struct {
	fileApi.AssetApi
}

type ArticleApi struct {
	articleApi.CommentApi
	articleApi.BulletinApi
	articleApi.ContentApi
}

type MessageApi struct {
	messageApi.CommentApi
}
