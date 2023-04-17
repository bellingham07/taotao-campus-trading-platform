package orderLogic

import (
	"com.xpwk/go-gin/model/response"
	orderRepository "com.xpwk/go-gin/repository/order"
	"github.com/gin-gonic/gin"
)

var InfoLogic = new(OrderInfoLogic)

type OrderInfoLogic struct {
}

func (*OrderInfoLogic) ListByUserId(userId int64) gin.H {
	orders := orderRepository.InfoRepository.ListByUserIdOrderByStatusDoneCreate(userId)
	if orders == nil {
		return gin.H{"code": response.FAIL, "msg": response.ERROR}
	}
	return gin.H{"code": response.OK, "msg": response.SUCCESS, "data": orders}
}
