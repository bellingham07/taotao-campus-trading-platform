package cmt

import (
	"context"
	"errors"
	"go-go-zero/service/cmdty/model/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"k8s.io/apimachinery/pkg/util/rand"
	"time"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByInfoIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListByInfoIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByInfoIdLogic {
	return &ListByInfoIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListByInfoIdLogic) ListByInfoId(req *types.CmdtyIdReq) ([]mongodb.CmdtyCmt, error) {
	var cmdtyId = req.Id

	pipeline1 := []bson.M{
		{
			"$match": bson.M{
				"cmdty_id": cmdtyId,
				"root_id":  0,
			},
		},
		{
			"$sort": bson.M{
				"create_at": -1, // -1 表示倒序
			},
		},
	}
	CmtsCursor, err := l.svcCtx.CmdtyCmt.Aggregate(l.ctx, pipeline1)
	if err != nil {
		logx.Infof("[MONGO ERROR] ListByInfoId 获取子评论错误 %v\n", err)
		return nil, errors.New("加载评论错误！")
	}
	var cmts []mongodb.CmdtyCmt
	if err = CmtsCursor.All(l.ctx, &cmts); err != nil {
		logx.Infof("[MONGO ERROR] ListByInfoId 获取子评论错误 %v\n", err)
		return nil, errors.New("加载评论错误！")
	}
	for idx, cmt := range cmts {
		pipeline2 := []bson.M{
			{
				"$match": bson.M{
					"cmdty_id": cmdtyId,
					"root_id":  cmt.Id,
				},
			},
			{
				"$sort": bson.M{
					"create_at": -1,
				},
			},
		}

		subCmtsCursor, err := l.svcCtx.CmdtyCmt.Aggregate(l.ctx, pipeline2)
		if err != nil {
			logx.Infof("[MONGO ERROR] ListByInfoId 获取子评论错误 %v\n", err)
			continue
		}
		var subCmt []mongodb.CmdtyCmt
		if err = subCmtsCursor.All(l.ctx, &subCmt); err != nil {
			logx.Infof("[MONGO ERROR] ListByInfoId 解析子评论错误 %v\n", err)
			continue
		}
		cmts[idx].SubCmt = subCmt
	}
	go l.deleteExpiredCmts(cmts)
	return cmts, nil
}

func (l *ListByInfoIdLogic) deleteExpiredCmts(cmts []mongodb.CmdtyCmt) {
	var (
		randNum = rand.Int63nRange(1, 99)
		ids     = make([]int64, 0)
	)
	if randNum > 5 {
		return
	}
	for _, cmt := range cmts {
		l.isExpired(cmt.Id, cmt.CreateAt, ids)
		for _, subCmt := range cmt.SubCmt {
			l.isExpired(subCmt.Id, subCmt.CreateAt, ids)
		}
	}
	if len(ids) > 0 {
		filter := bson.M{"_id": bson.M{"$in": ids}}
		_, err := l.svcCtx.CmdtyCmt.DeleteMany(l.ctx, filter)
		if err != nil {
			logx.Infof("[MONGO ERROR] deleteExpiredCmts 删除过期评论失败 %v\n", err)
		}
	}
}

func (l *ListByInfoIdLogic) isExpired(id int64, createTime time.Time, ids []int64) {
	var thirtyDaysBefore = time.Now().Add(-30 * 24 * time.Hour)
	if thirtyDaysBefore.After(createTime) {
		ids = append(ids, id)
	}
}
