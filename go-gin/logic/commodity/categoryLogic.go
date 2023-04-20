package commodityLogic

import (
	"com.xpwk/go-gin/cache"
	"com.xpwk/go-gin/model/response"
	commodityRepository "com.xpwk/go-gin/repository/commodity"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

var CategoryLogic = new(CommodityCategoryLogic)

type CommodityCategoryLogic struct {
}

func (*CommodityCategoryLogic) List() gin.H {
	categoriesStr1, err := cache.RedisClient.GET(cache.COMMODITYCATEGORY)
	if err != nil {
		return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": categoriesStr1}
	}
	categories := commodityRepository.CategoryRepository.QueryAll()
	categoriesStr2, err := json.Marshal(categories)
	if err != nil {
		return gin.H{"code": response.FAIL, "msg": response.ERROR, "data": categoriesStr1}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": categoriesStr2}
}
