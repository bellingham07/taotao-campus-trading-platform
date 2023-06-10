package tag

import (
	"context"
	"errors"
	jsoniter "github.com/json-iterator/go"
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

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (l *ListTagLogic) ListTag() (resp []*model.CmdtyTag, err error) {
	result, err := l.svcCtx.RedisClient.Get(l.ctx, utils.CmdtyTag).Result()
	if err != nil {
		logx.Debugf("[REDIS ERROR] ListTag redis中的 cmdtyTag 已被淘汰，请运营进行同步 " + err.Error())
		resp = l.svcCtx.CmdtyTag.List()
		if resp != nil {
			return nil, errors.New("出错啦！😢")
		}
		return resp, nil
	}
	resp = make([]*model.CmdtyTag, 0)
	json.Unmarshal([]byte(result), resp)
	return
}
