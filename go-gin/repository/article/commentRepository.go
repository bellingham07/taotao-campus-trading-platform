package articleRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
	"time"
)

var TagRepository = new(CommodityTagRepository)

type CommodityTagRepository struct {
}

func commodity_tag() string {
	return "commodity_tag"
}

func (*CommodityTagRepository) QueryAll() (cc []model.CommodityTag) {
	if err := repository.GetDB().Table(commodity_tag()).Find(&cc).Error; err != nil {
		return nil
	}
	return
}

func (*CommodityTagRepository) Insert(tag *model.CommodityTag) error {
	tag.UpdateAt = time.Now()
	tag.CreateAt = time.Now()
	if err := repository.GetDB().Table(commodity_tag()).Create(tag).Error; err != nil {
		return err
	}
	return nil
}

func (*CommodityTagRepository) DeleteById(id int64) error {
	var category = model.CommodityTag{Id: id}
	if err := repository.GetDB().Table(commodity_tag()).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}
