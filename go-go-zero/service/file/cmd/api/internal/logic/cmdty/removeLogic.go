package cmdty

import (
	"context"
	"errors"
	"go-go-zero/service/file/cmd/api/internal/logic"
	"go-go-zero/service/file/cmd/api/internal/types"
	"go-go-zero/service/file/model"
	"sync"
	"xorm.io/xorm"

	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/file/cmd/api/internal/svc"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.IdCmdtyIdReq) error {
	var (
		id    = req.IdReq.Id
		order int64
		flag  = false
		fcs   = make([]model.FileCmdty, 0)
	)

	err := l.svcCtx.FileCmdty.Where("`cmdty_id`= ?", req.Id, req.CmdtyId).Asc("`order`").Find(&fcs)
	if err != nil {
		return errors.New("找不到这张图片啦？😵")
	}

	for _, fc := range fcs {
		if fc.Id == id {
			flag = true
			order = fc.Order
			break
		}
	}
	if !flag {
		return errors.New("没有这张图片！😶")
	}

	_, err = l.svcCtx.Xorm.Transaction(func(s *xorm.Session) (interface{}, error) {
		var wg sync.WaitGroup

		tx := s.Table("file_cmdty")
		if _, err = tx.Delete(&model.FileCmdty{Id: fcs[0].Id}); err != nil {
			logx.Infof("[DB ERROR] Cmdty Remove %v", err)
			return nil, errors.New("删除失败！")
		}

		// 如果是第一张图片就是封面
		newOrders := make([]model.FileCmdty, 0)
		if order == 1 {
			for _, fc := range fcs[1:] {
				newOrder := model.FileCmdty{
					Id:    fc.Id,
					Order: fc.Order - 1,
				}
				newOrders = append(newOrders, newOrder)
			}
		} else { // 去掉下标为 order-1 的那张图片
			newOrders = append(newOrders, fcs[:order-1]...)
			for _, fc := range fcs[order:] {
				newOrder := model.FileCmdty{
					Id:    fc.Id,
					Order: fc.Order - 1,
				}
				newOrders = append(newOrders, newOrder)
			}
		}

		cmdtyservice.
		if _, err = tx.Update(&newOrders); err != nil {
			return nil, errors.New("删除失败！")
		}

		commonLogic := logic.NewCommonLogic(l.ctx, l.svcCtx)
		commonLogic.Delete()
	})

	return nil
}
