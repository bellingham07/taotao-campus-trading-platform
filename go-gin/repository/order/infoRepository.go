package orderRepository

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/repository"
	"log"
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
		log.Println("[DB ERROR]: ORDER CANCEL NOT FOUND")
		return nil
	}
	return info
}

func (*OrderInfoRepository) Insert(info *model.OrderInfo) error {
	if err := repository.GetDB().Table(order_info()).Create(info).Error; err != nil {
		return err
	}
	return nil
}

func (*OrderInfoRepository) UpdateById(info *model.OrderInfo) error {
	if err := repository.GetDB().Table(order_info()).Updates(info).Error; err != nil {
		log.Println("[DB ERROR]: ORDER CANCEL UPDATE FAIL")
		return err
	}
	return nil
}

func (*OrderInfoRepository) UpdateByIdAndStatus(info *model.OrderInfo, status int8) error {
	if err := repository.GetDB().Table(order_info()).Where("status = ?", status).Updates(info).Error; err != nil {
		log.Println("[DB ERROR]: ORDER CANCEL UPDATE FAIL")
		return err
	}
	return nil
}
