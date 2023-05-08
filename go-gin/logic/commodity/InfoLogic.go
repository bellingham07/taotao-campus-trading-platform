package commodityLogic

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/request"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/repository"
	commodityRepository "com.xpdj/go-gin/repository/commodity"
	"com.xpdj/go-gin/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var CommodityInfo = new(CommodityInfoLogic)

type CommodityInfoLogic struct {
}

// UpdateInfoAndArticle 更新草稿或者已发布，flag为标识（true为更新已发布，false为更新草稿）
func (il *CommodityInfoLogic) UpdateInfoAndArticle(draftDto *request.CommodityArticleDraft, isPublish bool) gin.H {
	info, articleContent := il.copyDraftAttribute(draftDto)
	if !isPublish {
		if err := il.update2(info, articleContent); err != nil {
			return response.GenH(response.FAIL, "操作失败，请重试！")
		}
		return response.GenH(response.OK, response.SUCCESS)
	}
	info.Status = 2
	info.PublishAt = time.Now()
	if err := il.update2(info, articleContent); err != nil {
		return response.GenH(response.FAIL, "操作失败，请重试！")
	}
	return response.GenH(response.OK, response.SUCCESS)
}

// SaveOrPublishInfoAndArticle 保存并发布商品信息，区分出售和购买
func (il *CommodityInfoLogic) SaveOrPublishInfoAndArticle(draftDto *request.CommodityArticleDraft, userId int64, cmdtyType int64, isPublish bool) interface{} {
	now := time.Now()
	infoDraft, articleContent := il.copyDraftAttribute(draftDto)
	infoDraft.CreateAt = now
	infoDraft.UserId = userId
	infoDraft.Type = cmdtyType
	articleContent.UserId = userId
	articleContent.CreateAt = now
	// 保存草稿
	if !isPublish {
		if err := il.create2(infoDraft, articleContent); err != nil {
			return response.GenH(response.FAIL, "操作失败，请重试！")
		}
		return response.GenH(response.OK, response.SUCCESS, gin.H{"id": draftDto.Id, "articleId": articleContent.Id})
	}
	// 直接发布
	infoDraft.Status = 2
	infoDraft.PublishAt = now
	if err := il.create2(infoDraft, articleContent); err != nil {
		return response.GenH(response.FAIL, "操作失败，请重试！", gin.H{"id": draftDto.Id, "articleId": articleContent.Id})
	}
	return response.GenH(response.OK, response.SUCCESS)
}

func (*CommodityInfoLogic) copyDraftAttribute(draftDto *request.CommodityArticleDraft) (*model.CommodityInfo, *model.ArticleContent) {
	articleContent := &model.ArticleContent{
		Title:    draftDto.Title,
		Content:  draftDto.Content,
		UpdateAt: time.Now(),
	}
	cmdtyInfo := &model.CommodityInfo{
		Name:  draftDto.Name,
		Model: draftDto.Model,
		Brand: draftDto.Brand,
		Price: draftDto.Price,
		Stock: draftDto.Stock,
		Tag:   draftDto.Tag,
	}
	return cmdtyInfo, articleContent
}

func (*CommodityInfoLogic) GetById(id int64, userId int64, exist bool) gin.H {
	var commodityInfo model.CommodityInfo
	key := utils.COMMODITYINFO + strconv.FormatInt(id, 10)
	commodityInfoMap, err := utils.RedisUtil.HGETALL(key)
	// 数据库也没有数据，防止缓存穿透
	if commodityInfoMap["id"] == "" {
		_ = utils.RedisUtil.HSET(key, commodityInfo)
		_ = utils.RedisUtil.EXPIRE(key, 30*time.Second)
		return response.GenH(response.FAIL, "没有此商品信息！")
	}
	// redis没有数据，就从数据库里查
	if err != nil {
		commodityInfo, err = commodityRepository.CommodityInfo.QueryById(id)
		// 数据无，设置空
		if err != nil {
			_ = utils.RedisUtil.HSET(key, commodityInfo)
			_ = utils.RedisUtil.EXPIRE(key, 30*time.Second)
			return response.GenH(response.FAIL, "没有此商品信息！")
		}
		// jwt中存在用户，判断是在访问自己的商品还是别人的
		if exist {
			// redis取出的值不为空则说明，redis中有
			if commodityInfo.UserId != userId {
				go HistoryLogic.UpdateHistory(id, userId)
			}
		}
		return response.GenH(response.OK, response.SUCCESS, commodityInfo)
	}
	// redis有数据
	if exist {
		// redis取出的值不为空则说明，redis中有
		if commodityInfo.UserId != userId {
			go HistoryLogic.UpdateHistory(id, userId)
		}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": commodityInfoMap}
}

func (*CommodityInfoLogic) RandomListByType(option int) gin.H {
	infos := commodityRepository.CommodityInfo.RandomListByType(option)
	if infos == nil {
		return gin.H{"code": response.FAIL, "msg": "系统繁忙，请稍后再试。"}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": infos}
}

func (il *CommodityInfoLogic) create2(cmdtyInfo *model.CommodityInfo, articleContent *model.ArticleContent) error {
	err := repository.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("article_content").Create(articleContent).Error; err != nil {
			return err
		}
		if err := tx.Table("commodity_info").Create(cmdtyInfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (il *CommodityInfoLogic) update2(cmdtyInfo *model.CommodityInfo, articleContent *model.ArticleContent) error {
	err := repository.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("article_content").Updates(articleContent).Error; err != nil {
			return err
		}
		if err := tx.Table("commodity_info").Updates(cmdtyInfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
