package room

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"go-go-zero/service/chat/model"
	"time"

	"go-go-zero/service/chat/cmd/api/internal/svc"
	"go-go-zero/service/chat/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoomLogic {
	return &CreateRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoomLogic) CreateRoom(req *types.CreateRoomReq) (*types.IdResp, error) {
	cr := &model.ChatRoom{
		Id:       idgen.NextId(),
		CmdtyId:  req.CmdtyId,
		SellerId: req.SellerId,
		Seller:   req.Seller,
		BuyerId:  req.BuyerId,
		Buyer:    req.Buyer,
		Cover:    req.Cover,
		CreateAt: time.Now(),
		Status:   req.CmdtyId,
	}
	_, err := l.svcCtx.Xorm.Table("chat_room").Insert(cr)
	resp := &types.IdResp{Id: cr.Id}
	if err == nil {
		return resp, nil
	}
	has, err := l.svcCtx.Xorm.Table("chat_room").
		Where("`cmdty_id` = ? AND `seller_id` = ? AND buyer_id` = ?",
			req.CmdtyId, req.SellerId, req.BuyerId).Get(cr)
	if has && err == nil {
		resp.Id = cr.Id
		return resp, err
	}
	return nil, errors.New("创建失败！😢")
}
