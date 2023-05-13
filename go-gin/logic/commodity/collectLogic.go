package commodityLogic

import (
	mqLogic "com.xpdj/go-gin/logic/rabbitmq"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/utils/cache"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

var CollectLogic = new(CommodityCollectLogic)

type CommodityCollectLogic struct {
}

func (*CommodityCollectLogic) Collect(id, userIdStr string) gin.H {
	key := cache.CommodityCollect + id
	err := cache.RedisUtil.SADD(key, userIdStr)
	if err == 0 {
		return response.ErrorMsg("不能再次收藏哦😊")
	}
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	go CollectUpdatePublisher(key, userId, true)
	return response.Ok()
}

func CollectUpdatePublisher(redisKey string, member int64, isCollect bool) {
	now := time.Now()
	ticker := time.NewTicker(time.Second * 30)
	message := &mqLogic.CcMessage{
		RedisKey:  redisKey,
		UserId:    member,
		Time:      now,
		IsCollect: isCollect,
	}
	body, _ := json.Marshal(message)
	publisher := mqLogic.CcPublisher()
	// 假如无法使用mq
	if publisher == nil {
		select {
		case <-ticker.C:
			if isCollect {
				go mqLogic.CollectCheckUpdate(message)
			} else {
				go mqLogic.CollectCheckDelete(message)
			}
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
			if isCollect {
				go mqLogic.CollectCheckUpdate(message)
			} else {
				go mqLogic.CollectCheckDelete(message)
			}
			return
		}
	}
}

func (*CommodityCollectLogic) Uncollect(idStr, userIdStr string) gin.H {
	key := cache.CommodityCollect + idStr
	isMember := cache.RedisUtil.SISMEMBER(key, userIdStr)
	if isMember {
		cache.RedisUtil.SREM(key, userIdStr)
		userId, _ := strconv.ParseInt(userIdStr, 10, 64)
		go CollectUpdatePublisher(key, userId, false)
		return response.Ok()
	}
	return response.ErrorMsg("你本来就没收藏人家嘛！😫")
}

func (*CommodityCollectLogic) List(userId string) gin.H {
	key := cache.CommodityCollect + userId
	ids := cache.RedisUtil.ZREVRANGE(key, 0, -1)
	if ids == nil {
		return response.Error()
	}
	var infosMap []map[string]string
	for _, id := range ids {
		key := cache.CommodityInfo + id
		infoMap := cache.RedisUtil.HGETALL(key)
		infosMap = append(infosMap, infoMap)
	}
	return response.OkData(infosMap)
}
