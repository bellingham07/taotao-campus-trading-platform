// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-go-zero/service/user/cmd/rpc/internal/logic"
	"go-go-zero/service/user/cmd/rpc/internal/svc"
	"go-go-zero/service/user/cmd/rpc/types"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) UpdateAvatar(ctx context.Context, in *__.AvatarReq) (*__.CodeResp, error) {
	l := logic.NewUpdateAvatarLogic(ctx, s.svcCtx)
	return l.UpdateAvatar(in)
}

func (s *UserServiceServer) RetrieveNameAndAvatar(ctx context.Context, in *__.IdReq) (*__.NameAvatarResp, error) {
	l := logic.NewRetrieveNameAndAvatarLogic(ctx, s.svcCtx)
	return l.RetrieveNameAndAvatar(in)
}
