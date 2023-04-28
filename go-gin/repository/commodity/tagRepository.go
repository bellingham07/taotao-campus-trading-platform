package commodityRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
	"time"
)

var TagRepository = new(CommodityTagRepository)

type CommodityTagRepository struct {
}

func (*CommodityTagRepository) tableName() string {
	return "commodity_tag"
}

func (ccr *CommodityTagRepository) QueryAll() (cc []model.CommodityTag) {
	if err := repository.GetDB().Table(ccr.tableName()).Find(&cc).Error; err != nil {
		return nil
	}
	return
}

func (ccr *CommodityTagRepository) Insert(tag *model.CommodityTag) error {
	tag.UpdateAt = time.Now()
	tag.CreateAt = time.Now()
	if err := repository.GetDB().Table(ccr.tableName()).Create(tag).Error; err != nil {
		return err
	}
	return nil
}

func (ccr *CommodityTagRepository) DeleteById(id int) error {
	var category = model.CommodityTag{Id: id}
	if err := repository.GetDB().Table(ccr.tableName()).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}
