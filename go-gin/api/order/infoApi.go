package orderApi

import (
	orderLogic "com.xpdj/go-gin/logic/order"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/request"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/yitter/idgenerator-go/idgen"
	"net/http"
	"strconv"
	"time"
)

type InfoApi struct {
}

func (*InfoApi) List(c *gin.Context) {
	userId := middleware.GetUserId(c)
	c.JSON(http.StatusOK, orderLogic.InfoLogic.ListByUserId(userId))
}

func (*InfoApi) GetById(c *gin.Context) {
	id := getOrderId(c)
	if id == 0 {
		return
	}
	c.JSON(http.StatusOK, orderLogic.InfoLogic.GetById(id))
}

func (*InfoApi) Buy(c *gin.Context) {
	orderDto := new(request.OrderDto)
	err := c.ShouldBind(orderDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	buyerId, name := middleware.GetUserIdAndName(c)
	order := &model.OrderInfo{
		Id:            idgen.NextId(),
		SellerId:      orderDto.OwnerId,
		SellerName:    orderDto.OwnerName,
		BuyerId:       buyerId,
		BuyerName:     name,
		CommodityId:   orderDto.CommodityId,
		CommodityName: orderDto.CommodityName,
		Location:      orderDto.Location,
		Payment:       orderDto.Payment,
		Status:        1,
		CreateAt:      time.Now(),
	}
	c.JSON(http.StatusOK, orderLogic.InfoLogic.SaveOrder(order))
}

func (*InfoApi) Sell(c *gin.Context) {
	orderDto := new(request.OrderDto)
	err := c.ShouldBind(orderDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	sellerId, name := middleware.GetUserIdAndName(c)
	order := &model.OrderInfo{
		Id:            idgen.NextId(),
		SellerId:      sellerId,
		SellerName:    name,
		BuyerId:       orderDto.OwnerId,
		BuyerName:     orderDto.OwnerName,
		CommodityId:   orderDto.CommodityId,
		CommodityName: orderDto.CommodityName,
		Location:      orderDto.Location,
		Payment:       orderDto.Payment,
		Status:        2,
		CreateAt:      time.Now(),
	}
	c.JSON(http.StatusOK, orderLogic.InfoLogic.SaveOrder(order))
}

func (*InfoApi) Cancel(c *gin.Context) {
	id := getOrderId(c)
	if id == 0 {
		return
	}
	sellerIdStr := c.PostForm("sellerId")
	userIdStr := middleware.GetUserIdStr(c)
	orderInfo := &model.OrderInfo{
		Id: id,
	}
	// 如果是卖家取消
	if sellerIdStr == userIdStr {
		orderInfo.Status = -2
	} else {
		orderInfo.Status = -1
	}
	c.JSON(http.StatusOK, orderLogic.InfoLogic.Cancel(orderInfo))
}

func (*InfoApi) Done(c *gin.Context) {
	id := getOrderId(c)
	if id == 0 {
		return
	}
	gb := c.Param("gb")
	if gb != "g" || gb != "b" {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return
	}
	orderInfo := &model.OrderInfo{
		Id:     id,
		DoneAt: time.Now(),
		Status: 3,
	}
	if gb == "g" {
		orderInfo.IsGood = "好👍"
	} else {
		orderInfo.IsGood = "差👎"
	}
	c.JSON(http.StatusOK, orderLogic.InfoLogic.UpdateOrderCheckStatus(orderInfo, 2))
}

func (*InfoApi) SellConfirm(c *gin.Context) {
	id := getOrderId(c)
	if id == 0 {
		return
	}
	userId := middleware.GetUserId(c)
	c.JSON(http.StatusOK, orderLogic.InfoLogic.SellConfirm(id, userId))
}

func getOrderId(c *gin.Context) int64 {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！"))
		return 0
	}
	return id
}
