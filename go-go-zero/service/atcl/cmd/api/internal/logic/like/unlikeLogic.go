package like

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/atcl/cmd/api/internal/types"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/atcl/cmd/api/internal/svc"
)

type UnlikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnlikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlikeLogic {
	return &UnlikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnlikeLogic) Unlike(req *types.IdReq) error {
	var (
		userId int64 = 408301323265285
		key          = utils.AtclLike + strconv.FormatInt(req.Id, 10)
	)
	_, err := l.svcCtx.Redis.SRem(l.ctx, key, userId).Result()
	if err != nil {
		logx.Infof("[REDIS ERROR] Unlike 文章取消点赞失败 %v\n")
		return errors.New("操作失败！😥")
	}
	return nil
}
