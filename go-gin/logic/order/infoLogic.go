package orderLogic

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	orderRepository "com.xpdj/go-gin/repository/order"
	"github.com/gin-gonic/gin"
)

var InfoLogic = new(OrderInfoLogic)

type OrderInfoLogic struct {
}

func (*OrderInfoLogic) ListByUserId(userId int64) gin.H {
	orders := orderRepository.InfoRepository.ListByUserIdOrderByStatusDoneCreate(userId)
	if orders == nil {
		return response.GenH(response.FAIL, response.ERROR)
	}
	return response.GenH(response.OK, response.SUCCESS, orders)
}

func (*OrderInfoLogic) GetById(id int64) interface{} {
	//key := utils.ORDERINFO + strconv.FormatInt(id, 10)
	//utils.RedisUtil.HGETALL()
	info := orderRepository.InfoRepository.QueryById(id)
	if info == nil {
		return response.GenH(response.FAIL, response.ERROR)
	}
	return response.GenH(response.OK, response.SUCCESS, info)
}

func (*OrderInfoLogic) SaveOrder(orderDto *model.OrderInfo) interface{} {
	err := orderRepository.InfoRepository.Insert(orderDto)
	if err == nil {
		return response.GenH(response.FAIL, "订单创建失败！😢请重试。")
	}
	return response.GenH(response.OK, response.SUCCESS)
}
