package api

type SystemApi struct {
	UserApi
	CommodityApi
	FileApi
	OrderApi
	ArticleApi
	MessageApi
}

var SystemApis = new(SystemApi)
