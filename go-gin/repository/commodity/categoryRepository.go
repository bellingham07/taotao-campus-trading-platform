package commodityRepository

import (
	"com.xpwk/go-gin/model"
	"com.xpwk/go-gin/repository"
)

var CategoryRepository = new(CommodityCategoryRepository)

type CommodityCategoryRepository struct {
}

func (*CommodityCategoryRepository) tableName() string {
	return "commodity_category"
}

func (ccr *CommodityCategoryRepository) QueryAll() (cc []model.CommodityCategory) {
	if err := repository.GetDB().Table(ccr.tableName()).Find(&cc).Error; err != nil {
		return nil
	}
	return
}
