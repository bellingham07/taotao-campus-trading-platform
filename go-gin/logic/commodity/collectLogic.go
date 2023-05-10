package commodityLogic

import (
	mqLogic "com.xpdj/go-gin/logic/mq"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/utils/cache"
	"com.xpdj/go-gin/utils/mq"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"time"
)

var CollectLogic = new(CommodityCollectLogic)

type CommodityCollectLogic struct {
}

func (*CommodityCollectLogic) Collect(id, userId string) gin.H {
	key := cache.COMMODITYCOLLECT + id
	err := cache.RedisUtil.SADD(key, userId)
	if err == 0 {
		return response.ErrorMsg("不能再次收藏哦😊")
	}
	go CollectUpdatePublisher(key, userId, true)
	return response.Ok()
}

func CollectUpdatePublisher(redisKey, member string, isCollect bool) {
	now := time.Now()
	ticker := time.NewTicker(time.Second * 30)
	message := mq.CcMessage{
		RedisKey:  redisKey,
		UserId:    member,
		Time:      now,
		IsCollect: isCollect,
	}
	body, _ := json.Marshal(message)
	publisher := mq.CcPublisher()
	if publisher == nil {
		select {
		case <-ticker.C:
			if isCollect {
				go mqLogic.CollectCheckUpdate(redisKey, member, now)
			} else {
				go mqLogic.CollectCheckDelete(redisKey, member)
			}
			return
		}
	}
	err := publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body:        body,
		})
	if err == nil {
		log.Println("发送成功咕咕咕咕咕咕过过过过过过过过过过过过过过过过过过过过过过")
	}
	if err != nil {
		log.Println("[RABBITMQ ERROR] ", err.Error())
		select {
		case <-ticker.C:
			if isCollect {
				go mqLogic.CollectCheckUpdate(redisKey, member, now)
			} else {
				go mqLogic.CollectCheckDelete(redisKey, member)
			}
			return
		}
	}
}

func (*CommodityCollectLogic) Uncollect(idStr, userIdStr string) gin.H {
	key := cache.COMMODITYCOLLECT + idStr
	isMember := cache.RedisUtil.SISMEMBER(key, userIdStr)
	if isMember {
		cache.RedisUtil.SREM(key, userIdStr)
		go CollectUpdatePublisher(key, userIdStr, false)
		return response.Ok()
	}
	return response.ErrorMsg("你本来就没收藏人家嘛！😫")
}

func (*CommodityCollectLogic) List(userId string) gin.H {
	key := cache.COMMODITYCOLLECT + userId
	ids := cache.RedisUtil.ZREVRANGE(key, 0, -1)
	if ids == nil {
		return response.Error()
	}
	var infosMap []map[string]string
	for _, id := range ids {
		key := cache.COMMODITYINFO + id
		infoMap, _ := cache.RedisUtil.HGETALL(key)
		infosMap = append(infosMap, infoMap)
	}
	return response.OkData(infosMap)
}
