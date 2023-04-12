package commodityRepository

import (
	"com.xpwk/go-gin/model"
)

var Commodity = new(CommodityRepository)

type CommodityRepository struct {
}

func (*CommodityRepository) ListOrderByTimeViewLike() []model.Commodity {

	return nil
}
