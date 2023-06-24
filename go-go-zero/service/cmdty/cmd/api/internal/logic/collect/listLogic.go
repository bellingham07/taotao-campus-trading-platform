package collect

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/model"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(userId int64) ([]model.CmdtyCollect, error) {
	var ccs = make([]model.CmdtyCollect, 0)
	err := l.svcCtx.CmdtyCollect.Where("user_id = ?", userId).Find(&ccs)
	if err != nil {
		return nil, errors.New("出错啦😫")
	}
	return ccs, nil
}
