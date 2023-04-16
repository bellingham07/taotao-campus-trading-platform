package commodityRepository

import (
	"com.xpwk/go-gin/model"
)

var CommodityInfo = new(CommodityInfoRepository)

type CommodityInfoRepository struct {
}

func (*CommodityInfoRepository) tableName() string {
	return "commodity_info"
}

func (*CommodityInfoRepository) ListOrderByTimeViewLike() []model.CommodityInfo {

	return nil
}
