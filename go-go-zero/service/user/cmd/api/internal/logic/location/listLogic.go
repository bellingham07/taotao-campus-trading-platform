package location

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/user/cmd/api/internal/svc"
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

func (l *ListLogic) List() ([]*model.UserLocation, error) {
	var uls = make([]*model.UserLocation, 0)
	result, err := l.svcCtx.Redis.Get(l.ctx, utils.UserLocation).Result()
	if err == nil {
		err = l.svcCtx.Json.Unmarshal([]byte(result), uls)
		if err == nil {
			return uls, nil
		}
		goto searchDB
	}
searchDB:
	err = l.svcCtx.UserLocation.Find(uls)
	if err != nil {
		return nil, errors.New("出错啦！😭")
	}
	return uls, nil
}
