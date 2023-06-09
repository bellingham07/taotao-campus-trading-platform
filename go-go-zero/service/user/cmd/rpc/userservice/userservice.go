// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"go-go-zero/service/user/cmd/rpc/types"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AvatarReq      = __.AvatarReq
	CodeResp       = __.CodeResp
	IdReq          = __.IdReq
	NameAvatarResp = __.NameAvatarResp

	UserService interface {
		UpdateAvatar(ctx context.Context, in *AvatarReq, opts ...grpc.CallOption) (*CodeResp, error)
		RetrieveNameAndAvatar(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*NameAvatarResp, error)
		IncrLike(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*CodeResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) UpdateAvatar(ctx context.Context, in *AvatarReq, opts ...grpc.CallOption) (*CodeResp, error) {
	client := __.NewUserServiceClient(m.cli.Conn())
	return client.UpdateAvatar(ctx, in, opts...)
}

func (m *defaultUserService) RetrieveNameAndAvatar(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*NameAvatarResp, error) {
	client := __.NewUserServiceClient(m.cli.Conn())
	return client.RetrieveNameAndAvatar(ctx, in, opts...)
}

func (m *defaultUserService) IncrLike(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*CodeResp, error) {
	client := __.NewUserServiceClient(m.cli.Conn())
	return client.IncrLike(ctx, in, opts...)
}
