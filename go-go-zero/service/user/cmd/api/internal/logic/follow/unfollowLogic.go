package follow

import (
	"context"
	"errors"
	"go-go-zero/service/user/cmd/api/internal/svc"
	"go-go-zero/service/user/cmd/api/internal/types"
	"go-go-zero/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnfollowLogic) Unfollow(req *types.IdReq) error {
	var (
		currentUser int64 = 408301323265285
		userId            = req.Id
	)
	uf := &model.UserFollow{
		UserId:       currentUser,
		FollowUserId: userId,
	}
	_, err := l.svcCtx.UserFollow.Delete(uf)
	if err != nil {
		return errors.New("操作失败啦！😢")
	}
	return nil
}
