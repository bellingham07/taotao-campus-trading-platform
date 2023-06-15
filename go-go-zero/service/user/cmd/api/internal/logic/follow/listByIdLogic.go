package follow

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/user/cmd/api/internal/svc"
	"go-go-zero/service/user/cmd/api/internal/types"
)

type ListByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByIdLogic {
	return &ListByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByIdLogic) ListById(req *types.PageReq) (interface{}, error) {
	var (
		userId int64 = 408301323265285
		offset       = (req.Page - 1) * 20
	)
	query := "select ui.`id`, ui.`name`, ui.`avatar`, ui.`intro` from `user_follow` uf left join `user_info` ui on uf.`follow_user_id` = ui.`id` where uf.user_id = ? order by uf.`create_at` desc limit ?, ?"
	uisMap, err := l.svcCtx.Xorm.QueryInterface(query, userId, 20, offset)
	if err != nil {
		return nil, errors.New("获取关注列表失败！😭")
	}
	return uisMap, nil
}
