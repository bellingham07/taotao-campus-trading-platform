package follow

import (
	"context"

	"go-go-zero/service/user/cmd/api/internal/svc"
	"go-go-zero/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *ListByIdLogic) ListById(req *types.IdReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
