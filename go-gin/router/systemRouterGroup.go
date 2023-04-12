package router

type SystemRouterGroup struct {
	UserRouter
	CommodityRouter
	OrderRouter
	ArticleRouter
	FileRouter
	MessageRouter
}
