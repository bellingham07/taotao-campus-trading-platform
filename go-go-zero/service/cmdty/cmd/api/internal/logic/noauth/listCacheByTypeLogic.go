package noauth

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/model"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCacheByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCacheByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCacheByTypeLogic {
	return &ListCacheByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCacheByTypeLogic) ListCacheByType(req *types.TypeReq) ([]*model.CmdtyInfo, error) {
	key := utils.CmdtySellingPrepared
	if req.Type == 2 {
		key = utils.CmdtyWantPrepared
	}
	cisMap, err := l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
	if err != nil {
		return nil, errors.New("出错啦！😥")
	}
	cis := make([]*model.CmdtyInfo, 0)
	for _, ciStr := range cisMap {
		ci := new(model.CmdtyInfo)
		err = l.svcCtx.Json.Unmarshal([]byte(ciStr), ci)
		if err != nil {
			logx.Infof("[JSON MARSHAL ERROR] ListCacheByType JSON反序列化失败 %v\n", err)
			continue
		}
		cis = append(cis, ci)
	}
	return cis, nil
}
