// Code generated by goctl. DO NOT EDIT.
// Source: cmdty.proto

package server

import (
	"context"

	"go-go-zero/service/cmdty/cmd/rpc/internal/logic"
	"go-go-zero/service/cmdty/cmd/rpc/internal/svc"
	"go-go-zero/service/cmdty/cmd/rpc/types"
)

type CmdtyServiceServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedCmdtyServiceServer
}

func NewCmdtyServiceServer(svcCtx *svc.ServiceContext) *CmdtyServiceServer {
	return &CmdtyServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CmdtyServiceServer) UpdateCover(ctx context.Context, in *__.CoverReq) (*__.CodeResp, error) {
	l := logic.NewUpdateCoverLogic(ctx, s.svcCtx)
	return l.UpdateCover(in)
}
