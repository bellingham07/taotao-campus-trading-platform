package follow

import (
	"context"
	"errors"
	"go-go-zero/service/user/model"
	"time"

	"go-go-zero/service/user/cmd/api/internal/svc"
	"go-go-zero/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.IdReq, currentUser int64) error {
	var userId = req.Id

	uf := &model.UserFollow{
		UserId:       currentUser,
		FollowUserId: userId,
		CreateAt:     time.Now(),
	}
	_, err := l.svcCtx.UserFollow.Insert(uf)
	if err != nil {
		return errors.New("操作失败啦！😢")
	}
	return nil
}
