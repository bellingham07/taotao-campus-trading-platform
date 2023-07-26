package noauth

import (
	"context"
	"errors"
	"fmt"
	"go-go-zero/service/cmdty/model"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdLogic {
	return &GetByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByIdLogic) GetById(req *types.IdReq) (map[string]interface{}, error) {
	var resp = make(map[string]interface{})
	ci := &model.CmdtyInfo{Id: req.Id}
	_, err := l.svcCtx.CmdtyInfo.Get(ci)
	if err != nil {
		fmt.Println("asdsdfhdlaskjf")
		return nil, errors.New("找不到这个商品")
	}
	resp["cmdtyInfo"] = ci
	resp["isCollected"] = 1
	return resp, nil
}
