package commodityLogic

import (
	mqLogic "com.xpdj/go-gin/logic/mq"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/repository"
	articleRepository "com.xpdj/go-gin/repository/article"
	commodityRepository "com.xpdj/go-gin/repository/commodity"
	"com.xpdj/go-gin/utils/cache"
	"com.xpdj/go-gin/utils/json"
	"com.xpdj/go-gin/utils/mq"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

var CommodityInfo = new(CommodityInfoLogic)

type CommodityInfoLogic struct {
}

// UpdateInfo 更新草稿或者已发布，flag为标识（true为更新已发布，false为更新草稿）
func (*CommodityInfoLogic) UpdateInfo(info *model.CommodityInfo, isPublish bool) gin.H {
	// 不是发布，即只更新内容
	if !isPublish {
		if err := commodityRepository.CommodityInfo.UpdateById(info); err != nil {
			return response.ErrorMsg("操作失败，请重试！")
		}
		return response.Ok()
	}
	// 是发布，就更新状态和内容
	info.Status = 2
	info.PublishAt = time.Now()
	if err := commodityRepository.CommodityInfo.UpdateById(info); err != nil {
		return response.ErrorMsg("操作失败，请重试！")
	}
	return response.Ok()
}

// SaveOrPublishInfo 保存并发布商品信息，区分出售和购买
func (*CommodityInfoLogic) SaveOrPublishInfo(infoDraft *model.CommodityInfo, userId int64, cmdtyType int64, isPublish bool) interface{} {
	now := time.Now()
	infoDraft.CreateAt = now
	infoDraft.UserId = userId
	infoDraft.Type = cmdtyType
	// 保存草稿
	if !isPublish {
		if err := commodityRepository.CommodityInfo.Insert(infoDraft); err != nil {
			return response.ErrorMsg("操作失败，请重试！")
		}
		return response.OkData(infoDraft.Id)
	}
	// 直接发布
	infoDraft.Status = 2
	infoDraft.PublishAt = now
	if err := commodityRepository.CommodityInfo.Insert(infoDraft); err != nil {
		return response.ErrorMsg("操作失败，请重试！")
	}
	return response.OkData(infoDraft.Id)
}

func (*CommodityInfoLogic) GetById(id int64, userId int64, isLogin bool) gin.H {
	idStr := strconv.FormatInt(id, 10)
	key := cache.CommodityInfo + idStr
	commodityInfoMap := cache.RedisUtil.HGETALL(key)
	// 数据库也没有数据，防止缓存穿透
	if commodityInfoMap["id"] == "" {
		_ = cache.RedisUtil.EXPIRE(key, 30*time.Second)
		return response.ErrorMsg("没有这个商品！你不要乱来呀😡")
	}
	collectKey := cache.CommodityCollect + idStr
	var isCollected int8 = 0
	var isMine int8 = 1
	// redis没有数据，就从数据库里查
	if commodityInfoMap == nil {
		commodityInfo := commodityRepository.CommodityInfo.QueryById(id)
		// 数据无，设置空
		if commodityInfo == nil {
			_ = cache.RedisUtil.HSET(key, map[string]string{"Id": ""})
			_ = cache.RedisUtil.EXPIRE(key, 30*time.Second)
			return response.ErrorMsg("没有这个商品！你不要乱来呀😡")
		}
		// 数据有
		go updateCommodityView(id)
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
			go updateCommodityView(id)
		}
	}
	return response.OkData(gin.H{"commodityInfo": commodityInfoMap, "isCollected": isCollected, "isMine": isMine})
}

func updateCommodityView(id int64) {
	key := cache.CommodityView + strconv.FormatInt(id, 10)
	// 1.如果hash的count字段设置成功，则说明，可以进行更新
	err := cache.RedisUtil.HSETNXPX(key, "count", 1, time.Minute*2)
	ticker := time.NewTicker(time.Second * 30)
	// 2.如果set失败，代表有点赞数，加1就好了
	if err != nil {
		_ = cache.RedisUtil.HINCRBY1(key, "count")
	}
	publisher := mq.VPublisher()
	// 3.假如无法使用mq
	if publisher == nil {
		select {
		case <-ticker.C:
			go mqLogic.ViewCheckUpdate(key, true)
			return
		}
	}
	vMessage := &mq.VMessage{
		RedisKey:    key,
		IsCommodity: true,
	}
	body, _ := json.Json.Marshal(vMessage)
	err = publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{
			ContentType: "appliction/json",
			Body:        body,
		})
	if err != nil {
		log.Println("[RABBITMQ ERROR] ", err.Error())
		select {
		case <-ticker.C:
			go mqLogic.ViewCheckUpdate(key, true)
			return
		}
	}
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
		if err := articleRepository.ArticleContent.Insert(articleContent); err != nil {
			return err
		}
		if err := commodityRepository.CommodityInfo.Insert(cmdtyInfo); err != nil {
			return err
		}
		return nil
	})
	return err
}
