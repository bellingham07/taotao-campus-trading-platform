package commodityLogic

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	commodityRepository "com.xpdj/go-gin/repository/commodity"
	"com.xpdj/go-gin/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

var CategoryLogic = new(CommodityCategoryLogic)

type CommodityCategoryLogic struct {
}

func (*CommodityCategoryLogic) List() gin.H {
	tagStr, err := utils.RedisUtil.GET(utils.COMMODITYCATEGORY)
	if err == nil {
		return response.GenH(response.OK, response.SUCCESS, tagStr)
	}
	tags := commodityRepository.TagRepository.QueryAll()
	if tags == nil {
		return response.GenH(response.FAIL, response.ERROR)
	}
	tagJsonStr, _ := json.Marshal(tags)
	return response.GenH(response.OK, response.SUCCESS, tagJsonStr)
}

func (*CommodityCategoryLogic) Add(tag *model.CommodityTag) gin.H {
	err := commodityRepository.TagRepository.Insert(tag)
	if err != nil {
		return gin.H{"code": response.FAIL, "msg": "添加出错，请重试。"}
	}
	tags := commodityRepository.TagRepository.QueryAll()
	if tags == nil {
		return response.GenH(response.FAIL, response.ERROR)
	}
	err = utils.RedisUtil.SET(utils.COMMODITYCATEGORY, tags, 0)
	if err != nil {
		log.Println("RemoveById 商品tag更新至redis出错！" + err.Error())
	}
	return response.GenH(response.OK, response.SUCCESS, tags)
}

func (*CommodityCategoryLogic) RemoveById(id int) gin.H {
	err := commodityRepository.TagRepository.DeleteById(id)
	if err != nil {
		return response.GenH(response.FAIL, "删除出错，请重试。")
	}
	tags := commodityRepository.TagRepository.QueryAll()
	if tags == nil {
		return response.GenH(response.FAIL, response.ERROR)
	}
	err = utils.RedisUtil.SET(utils.COMMODITYCATEGORY, tags, 0)
	if err != nil {
		log.Println("RemoveById 商品tag更新至redis出错！" + err.Error())
	}
	return response.GenH(response.OK, response.SUCCESS, tags)
}
