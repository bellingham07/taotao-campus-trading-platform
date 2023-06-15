package cmt

import (
	"context"
	"errors"
	"go-go-zero/service/cmdty/model/mongodb"
	"time"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

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
	if length := len(req.Content); length > 100 {
		return errors.New("评论太长了😭")
	}
	cc := &mongodb.CmdtyCmt{
		CmdtyId:  req.CmdtyId,
		UserId:   req.UserId,
		Content:  req.Content,
		RootId:   req.RootId,
		ToUserId: req.ToUserId,
		CreateAt: time.Now().Local(),
	}
	_, err := l.svcCtx.CmdtyCmt.InsertOne(l.ctx, cc)
	if err != nil {
		logx.Debugf("[MONGO ERROR] Cmt 评论插入mongodb失败 " + err.Error())
		return errors.New("评论失败😭")
	}
	return nil
}
