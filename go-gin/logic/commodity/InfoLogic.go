package commodityLogic

import (
	mqLogic "com.xpdj/go-gin/logic/rabbitmq"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/repository"
	articleRepository "com.xpdj/go-gin/repository/article"
	commodityRepository "com.xpdj/go-gin/repository/commodity"
	"com.xpdj/go-gin/utils/cache"
	"com.xpdj/go-gin/utils/json"
	"fmt"
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

func (*CommodityInfoLogic) GetById(id, userId int64) gin.H {
	idStr := strconv.FormatInt(id, 10)
	key := cache.CommodityInfo + idStr
	commodityInfoMap := cache.RedisUtil.HGETALL(key)
	// 数据库也没有数据，防止缓存穿透
	if commodityInfoMap["id"] == "nil" {
		_ = cache.RedisUtil.EXPIRE(key, 30*time.Second)
		return response.ErrorMsg("没有这个商品！你不要乱来呀😡1111")
	}
	collectKey := cache.CommodityCollect + idStr
	var collectFlag int8 = 0
	var isMine int8 = 1
	// redis没有数据，就从数据库里查
	if commodityInfoMap["id"] == "" {
		fmt.Println(1)
		commodityInfo := commodityRepository.CommodityInfo.QueryById(id)
		// 数据无，设置空
		if commodityInfo == nil {
			fmt.Println(2)

			_ = cache.RedisUtil.HSET(key, map[string]string{"id": "nil"})
			_ = cache.RedisUtil.EXPIRE(key, 30*time.Second)
			return response.ErrorMsg("没有这个商品！你不要乱来呀😡")
		}
		// 数据有
		go updateCommodityView(id)
		// jwt中存在用户，判断是在访问自己的商品还是别人的
		if userId != 0 {
			fmt.Println(3)

			// 当商品不是自己的时候更新足迹，并且把标识也做好（是否是自己的，是为1，就不需要收藏按钮，否则为0）
			if commodityInfo.UserId != userId {
				fmt.Println(4)

				collectFlag, isMine = updateHistoryAndIsCollected(id, userId, collectKey)
			}
		}
		return response.OkData(gin.H{"commodityInfo": commodityInfo, "isCollected": collectFlag, "isMine": isMine})
	}
	// redis有数据
	if userId != 0 {
		if commodityInfoMap["userId"] != strconv.FormatInt(userId, 10) {
			collectFlag, isMine = updateHistoryAndIsCollected(id, userId, collectKey)
			go updateCommodityView(id)
		}
	}
	return response.OkData(gin.H{"commodityInfo": commodityInfoMap, "isCollected": collectFlag, "isMine": isMine})
}

func updateCommodityView(id int64) {
	key := cache.CommodityView + strconv.FormatInt(id, 10)
	// 1.如果hash的count字段设置成功，则说明，可以进行更新
	err := cache.RedisUtil.HSETNX(key, "count", 1)
	ticker := time.NewTicker(time.Second * 30)
	// 2.如果set失败，代表有点赞数，加1就好了
	if err != nil {
		_ = cache.RedisUtil.HINCRBY1(key, "count")
	}
	publisher := mqLogic.VPublisher()
	vMessage := &mqLogic.VMessage{
		RedisKey:    key,
		IsCommodity: true,
	}
	// 3.假如无法使用mq
	if publisher == nil {
		select {
		case <-ticker.C:
			go mqLogic.ViewCheckUpdate(vMessage)
			return
		}
	}

	body, _ := json.Json.Marshal(vMessage)
	err = publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Println("[RABBITMQ ERROR] ", err.Error())
		select {
		case <-ticker.C:
			go mqLogic.ViewCheckUpdate(vMessage)
			return
		}
	}
}

func updateHistoryAndIsCollected(id, userId int64, collectKey string) (a1, a2 int8) {
	go HistoryLogic.UpdateHistory(id, userId)
	// 如果是别人的商品，就需要判断有没有收藏过
	if isMember := cache.RedisUtil.SISMEMBER(collectKey, userId); isMember {
		a1 = 1
	}
	return a1, 0
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

func (*CommodityInfoLogic) Like(id int64, userId int64) interface{} {
	key := cache.CommodityLike + strconv.FormatInt(id, 10)
	affect := cache.RedisUtil.SADD(key, userId)
	if affect == 0 {
		return response.ErrorMsg("不能重复点赞哦😊")
	}
	go LikeUpdatePublisher(key, userId)
	return response.Ok()
}

func (*CommodityInfoLogic) Unlike(id int64, userId int64) interface{} {
	key := cache.CommodityLike + strconv.FormatInt(id, 10)
	isMember := cache.RedisUtil.SISMEMBER(key, userId)
	if !isMember {
		return response.ErrorMsg("你本来就没有点赞🤡")
	}
	// 取消点赞，删去set中的member即可，没必要更改库
	go cache.RedisUtil.SREM(key, userId)
	return response.Ok()
}

func LikeUpdatePublisher(redisKey string, member int64) {
	now := time.Now()
	ticker := time.NewTicker(time.Second * 30)
	message := &mqLogic.LMessage{
		RedisKey:  redisKey,
		UserId:    member,
		Time:      now,
		IsArticle: false,
	}
	body, _ := json.Json.Marshal(message)
	publisher := mqLogic.LPublisher()
	// 假如无法使用mq
	if publisher == nil {
		select {
		case <-ticker.C:
			go mqLogic.LikeCheckUpdate(message)
			return
		}
	}
	err := publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Println("[RABBITMQ ERROR] ", err.Error())
		select {
		case <-ticker.C:
			go mqLogic.LikeCheckUpdate(message)
		}
	}
}
