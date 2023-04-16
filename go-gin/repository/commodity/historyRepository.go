package commodityRepository

import (
	"com.xpwk/go-gin/model"
	"com.xpwk/go-gin/repository"
)

var CommodityHistory = new(CommodityHistoryRepository)

type CommodityHistoryRepository struct {
}

func (*CommodityHistoryRepository) tableName() string {
	return "commodity_history"
}

func (*CommodityHistoryRepository) ListByUserId(userId int64) (commodityHistories []model.CommodityHistory) {
	if err := repository.GetDB().Table(CommodityHistory.tableName()).Find(&commodityHistories).Error; err != nil {
		return nil
	}
	return commodityHistories
}
