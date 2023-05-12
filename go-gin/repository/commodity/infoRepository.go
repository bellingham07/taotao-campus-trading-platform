package commodityRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
	"log"
)

var CommodityInfo = new(CommodityInfoRepository)

type CommodityInfoRepository struct {
}

func commodity_info() string {
	return "commodity_info"
}

func (*CommodityInfoRepository) ListOrderByTimeViewLike() []model.CommodityInfo {

	return nil
}

func (*CommodityInfoRepository) QueryById(id int64) *model.CommodityInfo {
	info := &model.CommodityInfo{Id: id}
	if err := repository.GetDB().Table(commodity_info()).First(info).Error; err != nil {
		return nil
	}
	return info
}

func (*CommodityInfoRepository) RandomListByType(option int) (infos []model.CommodityInfo) {
	if err := repository.GetDB().Table(commodity_info()).Where("type", option).Find(&infos).Limit(15).Error; err != nil {
		return nil
	}
	if err := repository.GetDB().Raw("select * from commodity_info where type = ? ORDER BY RAND() LIMIT 15", option); err != nil {
		return nil
	}
	return infos
}

func (*CommodityInfoRepository) Insert(info *model.CommodityInfo) error {
	if err := repository.GetDB().Table(commodity_info()).Create(info).Error; err != nil {
		log.Println("[GORM ERROR] CommodityInfo Insert Fail, Error: " + err.Error())
		return err
	}
	return nil
}

func (*CommodityInfoRepository) UpdateById(info *model.CommodityInfo) error {
	if err := repository.GetDB().Table(commodity_info()).Updates(info).Error; err != nil {
		log.Println("[GORM ERROR] CommodityInfo UpdateById Fail, Error: " + err.Error())
		return err
	}
	return nil
}

func (*CommodityInfoRepository) UpdateViewById(id, count int64) error {
	if err := repository.GetDB().Table(commodity_info()).Where("id = ?", id).Update("count = count + ?", count).Error; err != nil {
		log.Println("[GORM ERROR] CommodityInfo UpdateById Fail, Error: " + err.Error())
		return err
	}
	return nil
}
