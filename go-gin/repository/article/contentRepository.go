package articleRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
)

var ContentRepository = new(ArticleContentRepository)

type ArticleContentRepository struct {
}

func article_content() string {
	return "article_content"
}

func (*ArticleContentRepository) Insert(content *model.ArticleContent) error {
	if err := repository.GetDB().Table(article_content()).Create(content).Error; err != nil {
		return err
	}
	return nil
}

func (*ArticleContentRepository) QueryById(id int64) (info model.CommodityInfo, err error) {
	info.Id = id
	if err := repository.GetDB().Table(article_content()).First(&info).Error; err != nil {
		return info, err
	}
	return info, nil
}

func (*ArticleContentRepository) RandomListByType(option int) (infos []model.CommodityInfo) {
	if err := repository.GetDB().Table(article_content()).Where("type", option).Find(&infos).Limit(15).Error; err != nil {
		return nil
	}
	if err := repository.GetDB().Raw("select * from article_content where type = ? ORDER BY RAND() LIMIT 15", option); err != nil {
		return nil
	}
	return infos
}
