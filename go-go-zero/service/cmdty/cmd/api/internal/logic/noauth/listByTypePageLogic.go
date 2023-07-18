package noauth

import (
	"context"
	"errors"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"
	"go-go-zero/service/cmdty/model"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByTypePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByTypePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByTypePageLogic {
	return &ListByTypePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByTypePageLogic) ListByTypePage(req *types.TypePageReq) ([]model.CmdtyInfo, error) {
	var key = utils.CmdtySellNewest
	if req.Type == 2 {
		key = utils.CmdtyWantNewest
	}

	cis := l.fetchCache(key + strconv.FormatInt(int64(req.Page), 10))
	if len(cis) < 10 {
		cis = l.queryFromDB(req.Page, req.Type)
		if cis == nil {
			return nil, errors.New("出错啦😥")
		}
	}
	return cis, nil
}

func (l *ListByTypePageLogic) fetchCache(key string) []model.CmdtyInfo {
	var cisMap, err = l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
	if cisMap == nil || len(cisMap) == 0 || err != nil {
		logx.Infof("[REDIS ERROR] ListCache %v", err)
		return nil
	}

	cis := make([]model.CmdtyInfo, 0)
	for _, ciStr := range cisMap {
		var ci model.CmdtyInfo
		err = l.svcCtx.Json.Unmarshal([]byte(ciStr), &ci)
		if err != nil {
			continue
		}
		cis = append(cis, ci)
	}
	return cis
}

func (l *ListByTypePageLogic) queryFromDB(page int, ciType int8) []model.CmdtyInfo {
	var cis = make([]model.CmdtyInfo, 0)
	err := l.svcCtx.CmdtyInfo.Desc("publish_at").
		Where("`type` = ? AND `status` = ?", ciType, 2).
		Limit(20, page*20).Find(&cis)
	if err != nil {
		logx.Infof("[DB ERROR] queryFromDB %v", err)
		return nil
	}
	return cis
}
