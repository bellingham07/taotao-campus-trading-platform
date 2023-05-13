package articleRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
	"log"
)

var ArticleContent = new(ArticleContentRepository)

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

func (*ArticleContentRepository) UpdateById(content *model.ArticleContent) error {
	if err := repository.GetDB().Table(article_content()).Updates(content).Error; err != nil {
		return err
	}
	return nil
}

func (*ArticleContentRepository) QueryById(id int64) *model.CommodityInfo {
	info := &model.CommodityInfo{
		Id: id,
	}
	if err := repository.GetDB().Table(article_content()).First(&info).Error; err != nil {
		log.Println("[GORM ERROR] ArticleContent QueryById Fail, Error: " + err.Error())
		return nil
	}
	return info
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

func (*ArticleContentRepository) UpdateViewById(id, count int64) error {
	if err := repository.GetDB().Table(article_content()).Where("id = ?", id).Update("view = view + ?", count).Error; err != nil {
		log.Println("[GORM ERROR] ArticleContent UpdateViewById Fail, Error: " + err.Error())
		return err
	}
	return nil
}

func (*ArticleContentRepository) UpdateLikeById(id int64) error {
	log.Println(123123)
	if err := repository.GetDB().Raw("update ? set like = like + 1 where id = ?", article_content(), id).Error; err != nil {
		log.Println("[GORM ERROR] UserInfo UpdateLikeById Fail, Error: " + err.Error())
		return err
	}
	return nil
}
