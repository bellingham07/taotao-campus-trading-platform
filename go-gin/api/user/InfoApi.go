package userApi

import (
	userLogic "com.xpdj/go-gin/logic/user"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/request"
	"com.xpdj/go-gin/model/response"
	"com.xpdj/go-gin/router/middleware"
	"com.xpdj/go-gin/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type InfoApi struct {
}

func (*InfoApi) UserLogin(c *gin.Context) {
	var loginUser = new(request.LoginUser)
	_ = c.ShouldBind(loginUser)

	// TODO
	//if err != nil || loginUser.ValidCode == "" {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": response.FAIL,
	//		"msg":  "请输入正确验证码",
	//	})
	//	return
	//}
	c.JSON(http.StatusOK, userLogic.UserInfo.Login(loginUser))
}

func (*InfoApi) Logout(c *gin.Context) {
	userId := middleware.GetUserIdStr(c)
	key := utils.USERLOGIN + userId
	_ = utils.RedisUtil.DEL(key)
	c.JSON(http.StatusOK, response.GenH(response.OK, "期待下一次遇见！😊"))

}

func (*InfoApi) GetInfoById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.GenH(response.FAIL, "请求参数错误！"))
		return
	}
	c.JSON(http.StatusOK, userLogic.UserInfo.GetUserById(id))
}

func (*InfoApi) UpdateInfo(c *gin.Context) {
	info := new(model.UserInfo)
	if err := c.ShouldBind(info); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, response.GenH(response.FAIL, "请求参数错误！"))
		return
	}
	userId := middleware.GetUserId(c)
	info.Id = userId
	log.Printf("%+v\n", info)

	c.JSON(http.StatusOK, userLogic.UserInfo.UpdateInfo(info))
}

func (*InfoApi) Register(c *gin.Context) {
	// TODO
	var registerUser = new(request.RegisterUser)
	err := c.ShouldBind(registerUser)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, response.GenH(response.FAIL, "请求参数错误！"))
		return
	}
	c.JSON(http.StatusOK, userLogic.UserInfo.Register(registerUser))
}
