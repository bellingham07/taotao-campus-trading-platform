package commodityLogic

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/request"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/repository"
	articleRepository "com.xpdj/go-gin/repository/article"
	commodityRepository "com.xpdj/go-gin/repository/commodity"
	"com.xpdj/go-gin/utils/cache"
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
			return response.ErrorMsg("操作失败，请重试！")
		}
		return response.Ok()
	}
	info.Status = 2
	info.PublishAt = time.Now()
	if err := il.update2(info, articleContent); err != nil {
		return response.ErrorMsg("操作失败，请重试！")
	}
	return response.Ok()
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
			return response.ErrorMsg("操作失败，请重试！")
		}
		return response.OkData(gin.H{"id": draftDto.Id, "articleId": articleContent.Id})
	}
	// 直接发布
	infoDraft.Status = 2
	infoDraft.PublishAt = now
	if err := il.create2(infoDraft, articleContent); err != nil {
		return response.ErrorMsgData("操作失败，请重试！", gin.H{"id": draftDto.Id, "articleId": articleContent.Id})
	}
	return response.Ok()
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

func (*CommodityInfoLogic) GetById(id int64, userId int64, isLogin bool) gin.H {
	idStr := strconv.FormatInt(id, 10)
	key := cache.COMMODITYINFO + idStr
	commodityInfoMap, err := cache.RedisUtil.HGETALL(key)
	// 数据库也没有数据，防止缓存穿透
	if commodityInfoMap["id"] == "" {
		_ = cache.RedisUtil.EXPIRE(key, 30*time.Second)
		return response.ErrorMsg("没有这个商品！你不要乱来呀😡")
	}
	collectKey := cache.COMMODITYCOLLECT + idStr
	var isCollected int8 = 0
	var isMine int8 = 1
	// redis没有数据，就从数据库里查
	if err != nil {
		commodityInfo := commodityRepository.CommodityInfo.QueryById(id)
		// 数据无，设置空
		if commodityInfo == nil {
			_ = cache.RedisUtil.HSET(key, map[string]string{"Id": ""})
			_ = cache.RedisUtil.EXPIRE(key, 30*time.Second)
			return response.ErrorMsg("没有这个商品！你不要乱来呀😡")
		}
		// 数据有
		// jwt中存在用户，判断是在访问自己的商品还是别人的
		if isLogin {
			// 当商品不是自己的时候更新足迹
			if commodityInfo.UserId != userId {
				isCollected, isMine = updateHistoryAndPD(id, userId, collectKey)
			}
		}
		return response.OkData(gin.H{"commodityInfo": commodityInfo, "isCollected": isCollected, "isMine": isMine})
	}
	// redis有数据
	if isLogin {
		if commodityInfoMap["userId"] != strconv.FormatInt(userId, 10) {
			isCollected, isMine = updateHistoryAndPD(id, userId, collectKey)
		}
	}
	return response.OkData(gin.H{"commodityInfo": commodityInfoMap, "isCollected": isCollected, "isMine": isMine})
}

func updateHistoryAndPD(id, userId int64, collectKey string) (a1, a2 int8) {
	go HistoryLogic.UpdateHistory(id, userId)
	// 如果是别人的商品，就需要判断有没有收藏过
	if isMember := cache.RedisUtil.SISMEMBER(collectKey, userId); isMember {
		a1 = 1
	}
	a2 = 0
	return
}

func (*CommodityInfoLogic) RandomListByType(option int) gin.H {
	infos := commodityRepository.CommodityInfo.RandomListByType(option)
	if infos == nil {
		return response.OkMsg("系统繁忙，请稍后再试。")
	}
	return response.OkData(infos)
}

func (il *CommodityInfoLogic) create2(cmdtyInfo *model.CommodityInfo, articleContent *model.ArticleContent) error {
	err := repository.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := articleRepository.ContentRepository.Insert(articleContent); err != nil {
			return err
		}
		if err := commodityRepository.CommodityInfo.Insert(cmdtyInfo); err != nil {
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
