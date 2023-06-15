package model

import "time"

type (
	CmdtyInfo struct {
		Id        int64     `xorm:"id" json:"id"`                // id
		UserId    int64     `xorm:"user_id" json:"userId"`       // 用户id
		Cover     string    `xorm:"cover" json:"cover"`          // 封面图片
		Tag       string    `xorm:"tag" json:"tag"`              // 分类名
		Price     float64   `xorm:"price" json:"price"`          // 商品价格
		Brand     string    `xorm:"brand" json:"brand"`          // 品牌
		Model     string    `xorm:"model" json:"model"`          // 型号
		Intro     string    `xorm:"intro" json:"intro"`          // 商品介绍
		Old       string    `xorm:"old" json:"old"`              // 新旧程度
		Status    int64     `xorm:"status" json:"status"`        // 商品状态，默认1为草稿，2为发布，0为下架，-1为审核未通过需修改
		CreateAt  time.Time `xorm:"create_at" json:"createAt"`   // 创建时间
		PublishAt time.Time `xorm:"publish_at" json:"publishAt"` // 发布时间
		View      int64     `xorm:"view" json:"view"`            // 查看数量
		Collect   int64     `xorm:"collect" json:"collect"`      // 收藏数
		Type      int64     `xorm:"type" json:"type"`            // 1为售卖商品，2为收商品
		Like      int64     `xorm:"like" json:"like"`            // 点赞数
	}

	CmdtyDone struct {
		Id        int64     `xorm:"id" json:"id"`                // id
		UserId    int64     `xorm:"user_id" json:"userId"`       // 用户id
		Cover     string    `xorm:"cover" json:"cover"`          // 封面图片
		Tag       string    `xorm:"tag" json:"tag"`              // 分类名
		Price     float64   `xorm:"price" json:"price"`          // 商品价格
		Brand     string    `xorm:"brand" json:"brand"`          // 品牌
		Model     string    `xorm:"model" json:"model"`          // 型号
		Intro     string    `xorm:"intro" json:"intro"`          // 商品介绍
		PublishAt time.Time `xorm:"publish_at" json:"publishAt"` // 发布时间
		View      int64     `xorm:"view" json:"view"`            // 查看数量
		Collect   int64     `xorm:"collect" json:"collect"`      // 收藏数
		Type      int64     `xorm:"type" json:"type"`            // 1为售卖商品，2为收商品
		Like      int64     `xorm:"like" json:"like"`            // 点赞数
	}

	CmdtyCollect struct {
		Id       int64     `xorm:"id" json:"id"`            // id
		UserId   int64     `xorm:"user_id" json:"userId"`   // 用户id
		CmdtyId  int64     `xorm:"cmdty_id" json:"cmdtyId"` // 商品id
		Intro    string    `xorm:"intro" json:"intro"`      // 20字的简介
		Cover    string    `xorm:"cover" json:"cover"`
		Price    float64   `xorm:"price" json:"price"`
		Status   int64     `xorm:"status" json:"status"`      // 1存在，0失效
		CreateAt time.Time `xorm:"create_at" json:"createAt"` // 创建时间
	}

	CmdtyTag struct {
		Id       int64     `xorm:"id" json:"id"`              // 分类ID编号
		Name     string    `xorm:"name" json:"name"`          // 分类名称
		CreateBy int64     `xorm:"create_by" json:"createBy"` // 管理员的id
		CreateAt time.Time `xorm:"create_at" json:"createAt"` // 创建时间
		UpdateBy int64     `xorm:"update_by" json:"updateBy"` // 管理员的id
		UpdateAt time.Time `xorm:"update_at" json:"updateAt"` // 更新时间
	}
)
