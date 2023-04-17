package orderRepository

import (
	"com.xpwk/go-gin/model"
	"com.xpwk/go-gin/repository"
)

var InfoRepository = new(OrderInfoRepository)

type OrderInfoRepository struct {
}

func (*OrderInfoRepository) tableName() string {
	return "order_info"
}

func (*OrderInfoRepository) ListByUserIdOrderByStatusDoneCreate(userId int64) (orderInfos []model.OrderInfo) {
	if err := repository.GetDB().Table(InfoRepository.tableName()).Where("user_id = ?", userId).Find(&orderInfos); err != nil {
		return nil
	}
	return orderInfos
}
