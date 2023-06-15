package cmt

import (
	"context"
	"errors"
	"go-go-zero/service/cmdty/cmd/api/internal/types"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

type RemoveCmtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveCmtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveCmtLogic {
	return &RemoveCmtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCmtLogic) RemoveCmt(req *types.IdReq) error {
	_, err := l.svcCtx.CmdtyCmt.DeleteOne(l.ctx, bson.M{"_id": req.Id})
	if err != nil {
		return errors.New("操作失败😢")
	}
	return nil
}
