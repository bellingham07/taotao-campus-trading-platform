package model

// CommentCommodity 商品评论
type CommentCommodity struct {
	Id          int64  `json:"id"`           // bigint自增，id
	CommodityId string `json:"commodity_id"` // bigint自增
	RootId      int64  `json:"root_id"`      // 根留言
	ToUserId    int64  `json:"to_user_id"`   // 雪花算法，回复给到用户
	UserId      int64  `json:"user_id"`      // 雪花算法，留言的用户id
	Content     string `json:"content"`      // 留言内容
	CreateAt    string `json:"create_at"`    // 评论时间
}

// CommentArticle 文章评论
type CommentArticle struct {
	Id        int64  `json:"id"`           // bigint自增，id
	ArticleId string `json:"commodity_id"` // bigint自增
	RootId    int64  `json:"root_id"`      // 根留言
	ToUserId  int64  `json:"to_user_id"`   // 雪花算法，回复给到用户
	UserId    int64  `json:"user_id"`      // 雪花算法，留言的用户id
	Content   string `json:"content"`      // 留言内容
	CreateAt  string `json:"create_at"`    // 评论时间
}
