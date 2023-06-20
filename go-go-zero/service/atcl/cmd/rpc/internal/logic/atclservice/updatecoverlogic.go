package atclservicelogic

import (
	"context"

	"go-go-zero/service/atcl/cmd/rpc/internal/svc"
	"go-go-zero/service/atcl/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCoverLogic {
	return &UpdateCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCoverLogic) UpdateCover(in *__.CoverReq) (*__.CodeResp, error) {
	// todo: add your logic here and delete this line

	return &__.CodeResp{}, nil
}
