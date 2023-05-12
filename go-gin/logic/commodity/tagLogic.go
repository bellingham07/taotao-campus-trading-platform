package commodityLogic

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	commodityRepository "com.xpdj/go-gin/repository/commodity"
	"com.xpdj/go-gin/utils/cache"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

var CategoryLogic = new(CommodityCategoryLogic)

type CommodityCategoryLogic struct {
}

func (*CommodityCategoryLogic) List() gin.H {
	tagStr, err := cache.RedisUtil.GET(cache.CommodityCategory)
	if err == nil {
		return response.OkData(tagStr)
	}
	tags := commodityRepository.TagRepository.QueryAll()
	if tags == nil {
		return response.Error()
	}
	tagJsonStr, _ := json.Marshal(tags)
	return response.OkData(tagJsonStr)
}

func (*CommodityCategoryLogic) Add(tag *model.CommodityTag) gin.H {
	err := commodityRepository.TagRepository.Insert(tag)
	if err != nil {
		return gin.H{"code": response.ERROR, "msg": "添加出错，请重试。"}
	}
	tags := commodityRepository.TagRepository.QueryAll()
	if tags == nil {
		return response.Error()
	}
	err = cache.RedisUtil.SET2JSON(cache.CommodityCategory, tags, 0)
	if err != nil {
		log.Println("RemoveById 商品tag更新至redis出错！" + err.Error())
	}
	return response.OkData(tags)
}

func (*CommodityCategoryLogic) RemoveById(id int64) gin.H {
	err := commodityRepository.TagRepository.DeleteById(id)
	if err != nil {
		return response.ErrorMsg("删除出错，请重试。")
	}
	tags := commodityRepository.TagRepository.QueryAll()
	if tags == nil {
		return response.Error()
	}
	err = cache.RedisUtil.SET2JSON(cache.CommodityCategory, tags, 0)
	if err != nil {
		log.Println("RemoveById 商品tag更新至redis出错！" + err.Error())
	}
	return response.OkData(tags)
}
