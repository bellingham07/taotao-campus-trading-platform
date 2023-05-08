package orderRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
)

var InfoRepository = new(OrderInfoRepository)

type OrderInfoRepository struct {
}

func order_info() string {
	return "order_info"
}

func (*OrderInfoRepository) ListByUserIdOrderByStatusDoneCreate(userId int64) (orderInfos []model.OrderInfo) {
	if err := repository.GetDB().Table(order_info()).Where("user_id = ?", userId).Find(&orderInfos); err != nil {
		return nil
	}
	return orderInfos
}

func (*OrderInfoRepository) QueryById(id int64) *model.OrderInfo {
	info := &model.OrderInfo{
		Id: id,
	}
	if err := repository.GetDB().Table(order_info()).First(info); err != nil {
		return nil
	}
	return info
}
