package commodityRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
	"log"
)

var CommodityCollect = new(CommodityCollectRepository)

type CommodityCollectRepository struct {
}

func commodity_collect() string {
	return "commodity_collect"
}

func (*CommodityCollectRepository) Insert(collect *model.CommodityCollect) error {
	if err := repository.GetDB().Table(commodity_collect()).Create(collect).Error; err != nil {
		log.Println("[GORM ERROR] CollectCHeckAndUpdate Fail ", err.Error())
		return err
	}
	return nil
}

func (r *CommodityCollectRepository) DeleteByCmdtyIdAndUserId(cmdtyId, userId int64) error {
	if err := repository.GetDB().Table(commodity_collect()).
		Where("commodity_id = ? AND user_id = ?", cmdtyId, userId).
		Delete(&model.CommodityCollect{}).Error; err != nil {
		log.Println("[GORM ERROR] DeleteByCmdtyIdAndUserId Fail ", err.Error())
		return err
	}
	return nil
}
