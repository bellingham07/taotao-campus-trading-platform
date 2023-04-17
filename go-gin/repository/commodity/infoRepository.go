package commodityRepository

import (
	"com.xpwk/go-gin/model"
	"com.xpwk/go-gin/repository"
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

func (*CommodityInfoRepository) QueryById(id int64) (commodityInfo model.CommodityInfo, err error) {
	commodityInfo.Id = id
	if err := repository.GetDB().First(&commodityInfo).Error; err != nil {
		return commodityInfo, err
	}
	return commodityInfo, nil
}
