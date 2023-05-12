package request

import "time"

type CommodityArticleDraft struct {
	Id        int64     `json:"id"`        // bigint自增，id
	Name      string    `json:"name"`      // 名称
	UserId    int64     `json:"userId"`    // uuid，用户id
	Brand     string    `json:"brand"`     // 品牌
	Model     string    `json:"model"`     // 型号
	Price     float64   `json:"price"`     // 商品价格
	Status    int64     `json:"status"`    // 商品状态，默认1为发布，0为下架，-1为审核未通过需修改
	Stock     int64     `json:"stock"`     // 库存
	Tag       string    `json:"tag"`       // 分类名
	Type      int64     `json:"type"`      // 1为售卖商品，2为收商品
	CreateAt  time.Time `json:"createAt"`  // datetime，创建时间
	PublishAt time.Time `json:"publishAt"` // datetime，发布时间
	UpdateAt  time.Time `json:"updateAt"`  // datetime，更新时间
	Title     string    `json:"title"`
	Content   string    `json:"content"` // 帖子内容
}
