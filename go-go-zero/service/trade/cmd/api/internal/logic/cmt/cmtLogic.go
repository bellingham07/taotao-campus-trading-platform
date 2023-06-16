package cmt

import (
	"context"
	"errors"
	"go-go-zero/service/trade/model"
	"go-go-zero/service/user/cmd/rpc/userservice"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CmtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CmtLogic {
	return &CmtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmtLogic) Cmt(req *types.CmtReq) error {
	var tc = &model.TradeCmt{
		TradeId:  req.TradeId,
		UserId:   req.UserId,
		ToUserId: req.ToUserId,
		Content:  req.Content,
		Type:     req.Type,
		CreateAt: time.Now().Local(),
	}
	result, err := l.svcCtx.TradeCmt.InsertOne(l.ctx, tc)
	if result.InsertedID != 0 && err != nil {
		logx.Infof("[MONGO ERROR] Cmt 插入评论失败 %v\n", err)
		return errors.New("评论失败，请重试😢")
	}
	go l.updateNameAndAvatar(result.InsertedID)
	return nil
}

func (l *CmtLogic) updateNameAndAvatar(insertedID interface{}) {
	// 1 先从rpc服务获取 昵称 和 头像
	var id = insertedID.(int64)
	resp, err := l.svcCtx.UserRpc.RetrieveNameAndAvatar(l.ctx, &userservice.IdReq{Id: id})
	if resp.GetCode() == -1 || err != nil {
		logx.Infof("[RPC ERROR] updateNameAndAvatar 远程获取用户名称和头像错误，userId：%v\n", id)
		return
	}
	// 2 再更新mongo中的记录
	name := resp.Name
	avatar := resp.Avatar
	update := bson.M{"user": name, "userAvatar": avatar}
	result, err := l.svcCtx.TradeCmt.UpdateByID(l.ctx, insertedID, update)

	mc := result.ModifiedCount
	if mc <= 0 || err != nil {
		logx.Infof("[MONGO ERROR] updateNameAndAvatar 更新用户名称和头像失败 %v\n", err)
	}
}
