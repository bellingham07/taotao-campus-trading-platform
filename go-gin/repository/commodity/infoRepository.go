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

func (*CommodityInfoRepository) QueryById(id int64) (info model.CommodityInfo, err error) {
	info.Id = id
	if err := repository.GetDB().First(&info).Error; err != nil {
		return info, err
	}
	return info, nil
}

func (*CommodityInfoRepository) RandomListByType(option int) (infos []model.CommodityInfo) {
	if err := repository.GetDB().Where("type", option).Find(&infos).Limit(15).Error; err != nil {
		return nil
	}
	if err := repository.GetDB().Raw("select * from commodity_info where type = ? ORDER BY RAND() LIMIT 15", option); err != nil {
		return nil
	}
	return infos
}
