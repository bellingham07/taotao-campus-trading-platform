package router

type SystemGroup struct {
	UserRouter
	CommodityRouter
	OrderRouter
	ArticleRouter
	FileRouter
	MessageRouter
}
