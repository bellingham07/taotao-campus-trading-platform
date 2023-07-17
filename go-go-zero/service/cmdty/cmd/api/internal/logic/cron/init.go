package cron

import "go-go-zero/service/cmdty/cmd/api/internal/svc"

func InitCronJob(svcCtx *svc.ServiceContext) {
	ci2CacheLogic(svcCtx)
}
