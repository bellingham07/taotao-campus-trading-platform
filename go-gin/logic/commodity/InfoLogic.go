package commodityLogic

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	commodityRepository "com.xpdj/go-gin/repository/commodity"
	"com.xpdj/go-gin/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var CommodityInfo = new(CommodityInfoLogic)

type CommodityInfoLogic struct {
}

func (*CommodityInfoLogic) SaveCommodity() gin.H {
	return nil
}

func (*CommodityInfoLogic) GetById(id int64, userId int64, exist bool) gin.H {
	var commodityInfo model.CommodityInfo
	key := utils.COMMODITYINFO + strconv.FormatInt(id, 10)
	commodityInfoMap, err := utils.RedisUtil.HGETALL(key)
	// 数据库也没有数据，防止缓存穿透
	if commodityInfoMap["id"] == "" {
		_ = utils.RedisUtil.HSET(key, commodityInfo)
		_ = utils.RedisUtil.EXPIRE(key, 30*time.Second)
		return gin.H{"code": response.FAIL, "msg": "没有此商品信息"}
	}
	// redis没有数据，就从数据库里查
	if err != nil {
		commodityInfo, err = commodityRepository.CommodityInfo.QueryById(id)
		// 数据无，设置空
		if err != nil {
			_ = utils.RedisUtil.HSET(key, commodityInfo)
			_ = utils.RedisUtil.EXPIRE(key, 30*time.Second)
			return gin.H{"code": response.FAIL, "msg": "没有此商品信息"}
		}
		// jwt中存在用户，判断是在访问自己的商品还是别人的
		if exist {
			// redis取出的值不为空则说明，redis中有
			if commodityInfo.UserId != userId {
				go HistoryLogic.UpdateHistory(id, userId)
			}
		}
		return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": commodityInfo}
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
