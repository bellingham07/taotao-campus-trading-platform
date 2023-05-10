package orderLogic

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/repository"
	orderRepository "com.xpdj/go-gin/repository/order"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var InfoLogic = new(OrderInfoLogic)

type OrderInfoLogic struct {
}

func (*OrderInfoLogic) ListByUserId(userId int64) gin.H {
	orders := orderRepository.InfoRepository.ListByUserIdOrderByStatusDoneCreate(userId)
	if orders == nil {
		return response.Error()
	}
	return response.OkData(orders)
}

func (*OrderInfoLogic) GetById(id int64) interface{} {
	//key := utils.ORDERINFO + strconv.FormatInt(id, 10)
	//utils.RedisUtil.HGETALL()
	info := orderRepository.InfoRepository.QueryById(id)
	if info == nil {
		return response.Error()
	}
	return response.OkData(info)
}

func (*OrderInfoLogic) SaveOrder(orderDto *model.OrderInfo) gin.H {
	err := orderRepository.InfoRepository.Insert(orderDto)
	if err == nil {
		return response.ErrorMsg("订单创建失败！😢请重试。")
	}
	return response.Ok()
}

func (*OrderInfoLogic) UpdateOrder(orderInfo *model.OrderInfo) gin.H {
	err := orderRepository.InfoRepository.UpdateById(orderInfo)
	if err == nil {
		return response.ErrorMsg("操作失败！😢请重试。")
	}
	return response.Ok()
}

func (il *OrderInfoLogic) Cancel(orderInfo *model.OrderInfo) gin.H {
	h := response.Ok()
	_ = repository.GetDB().Transaction(func(tx *gorm.DB) error {
		info := orderRepository.InfoRepository.QueryById(orderInfo.Id)
		if info != nil {
			h = response.ErrorMsg("没有该订单信息！")
			return errors.New("[DB FAIL]: ORDER CANCEL NOT FOUND")
		}
		if info.Status == -1 || info.Status == -2 {
			h = response.ErrorMsg("订单已经被取消！")
			return nil
		}
		if info.Status == 3 {
			h = response.ErrorMsg("订单已经完成，不能取消！")
			return nil
		}
		if h = il.UpdateOrder(orderInfo); h["code"] == response.ERROR {
			h = response.ErrorMsg("操作失败！😢请重试。")
			return errors.New("[DB FAIL]: ORDER CANCEL UPDATE ERROR")
		}
		return nil
	})
	return h
}

func (*OrderInfoLogic) SellConfirm(id int64, userId int64) interface{} {
	err := repository.GetDB().Table("order_info").Where("id = ?, sellerId = ?, status = 1", id, userId).Update("status = ?", 2).Error
	if err != nil {
		return response.ErrorMsg("你不能确认这个订单！")
	}
	return response.Ok()
}

func (*OrderInfoLogic) UpdateOrderCheckStatus(orderInfo *model.OrderInfo, status int8) gin.H {
	err := orderRepository.InfoRepository.UpdateByIdAndStatus(orderInfo, status)
	if err == nil {
		return response.ErrorMsg("操作失败！😢请重试。")
	}
	return response.Ok()
}
