// Code generated by goctl. DO NOT EDIT.
package types

type IdReq struct {
	Id int64 `path:"id" json:"id"`
}

type TypeReq struct {
	Type int64 `path:"type" json:"type"`
}

type IdDoneTypeReq struct {
	Id   int64 `form:"id" path:"id" json:"id"`
	Done int64 `form:"done" path:"done" json:"done"`
	Type int64 `form:"type" json:"type" path:"type"`
}

type CmdtyIdReq struct {
	Id int64 `path:"cmdtyId" json:"cmdtyId"`
}

type InfoReq struct {
	Id        int64   `json:"id"`        // id
	UserId    int64   `json:"userId"`    // 用户id
	Cover     string  `json:"cover"`     // 封面图片
	Tag       string  `json:"tag"`       // 分类名
	Price     float64 `json:"price"`     // 商品价格
	Brand     string  `json:"brand"`     // 品牌
	Model     string  `json:"model"`     // 型号
	Intro     string  `json:"intro"`     // 商品介绍
	Status    int64   `json:"status"`    // 商品状态，默认1为草稿，2为发布，0为下架，-1为审核未通过需修改
	CreateAt  string  `json:"createAt"`  // 创建时间
	PublishAt string  `json:"publishAt"` // 发布时间
	Type      int64   `json:"type"`      // 1为售卖商品，2为收商品
}

type IdsReq struct {
	Ids []int64 `json:"ids"`
}

type CmtReq struct {
	CmdtyId  int64  `json:"cmdtyId"`  // 对应的商品id
	UserId   int64  `json:"userId"`   // 留言的用户id
	Content  string `json:"content"`  // 留言内容
	RootId   int64  `json:"rootId"`   // 根留言
	ToUserId int64  `json:"toUserId"` // 回复给到用户
}

type InfoResp struct {
	Id        int64   `json:"id"`     // id
	UserId    int64   `json:"userId"` // 用户id
	Cover     string  `json:"cover"`  // 封面图片
	Tag       string  `json:"tag"`    // 分类名
	Price     float64 `json:"price"`  // 商品价格
	Brand     string  `json:"brand"`  // 品牌
	Model     string  `json:"model"`  // 型号
	Intro     string  `json:"intro"`  // 商品介绍
	Old       string  `json:"old"`
	Status    int64   `json:"status"`    // 商品状态，默认1为草稿，2为发布，0为下架，-1为审核未通过需修改
	CreateAt  string  `json:"createAt"`  // 创建时间
	PublishAt string  `json:"publishAt"` // 发布时间
	View      int64   `json:"view"`      // 查看数量
	Collect   int64   `json:"collect"`   // 收藏数
	Type      int64   `json:"type"`      // 1为售卖商品，2为收商品
	Like      int64   `json:"like"`
}

type InfoLiteResp struct {
	Id     int64   `json:"id"`     // id
	UserId int64   `json:"userId"` // 用户id
	Cover  string  `json:"cover"`
	Price  float64 `json:"price"` // 商品价格
	Intro  string  `json:"intro"`
}

type CollectResp struct {
	Id       int64   `json:"id"`
	UserId   int64   `json:"userId"`
	CmdtyId  int64   `json:"cmdtyId"`
	Cover    string  `json:"cover"`
	Price    float64 `json:"price"`
	Status   int64   `json:"status"`
	CreateAt string  `json:"createAt"`
}

type TagResp struct {
	Id       int64  `json:"id"`        // 分类ID编号
	Name     string `json:"name"`      // 分类名称
	CreateBy int64  `json:"createBy""` // 管理员的id
	CreateAt string `json:"createAt"`  // 创建时间
	UpdateBy int64  `json:"updateBy"`  // 管理员的id
	UpdateAt string `json:"updateAt"`  // 更新时间
}

type HistoryResp struct {
	Id    int64   `json:"id"`
	Cover string  `json:"cover"`
	Price float64 `json:"price"`
}

type CmtResp struct {
	Id       int64  `json:"id"`       // id
	CmdtyId  int64  `json:"cmdtyId"`  // 对应的商品id
	UserId   int64  `json:"userId"`   // 留言的用户id
	Content  string `json:"content"`  // 留言内容
	RootId   int64  `json:"rootId"`   // 根留言
	ToUserId int64  `json:"toUserId"` // 回复给到用户
	CreateAt string `json:"createAt"` // 评论时间
}
