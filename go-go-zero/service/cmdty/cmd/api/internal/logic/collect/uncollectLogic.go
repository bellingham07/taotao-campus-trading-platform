package collect

import (
	"context"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UncollectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUncollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UncollectLogic {
	return &UncollectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UncollectLogic) Uncollect(req *types.IdReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
