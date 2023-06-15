package tag

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/model"

	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

type ListTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTagLogic {
	return &ListTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTagLogic) ListTag() ([]*model.CmdtyTag, error) {
	var cts = make([]*model.CmdtyTag, 0)
	ctsStr, err := l.svcCtx.Redis.Get(l.ctx, utils.CmdtyTag).Result()
	if err == nil && ctsStr != "" {
		err = l.svcCtx.Json.Unmarshal([]byte(ctsStr), cts)
		if err != nil {
			return nil, errors.New("出错啦！😢")
		}
		return cts, nil
	}
	logx.Debugf("[REDIS ERROR] ListTag redis中的 cmdtyTag 已被淘汰，请运营进行同步 %v\n", err)
	if err = l.svcCtx.CmdtyTag.Find(cts); err != nil {
		return nil, errors.New("出错啦！😢")
	}
	return cts, nil
}
