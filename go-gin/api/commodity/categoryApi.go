package commodityApi

import (
	commodityLogic "com.xpwk/go-gin/logic/commodity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryApi struct {
}

func (*CategoryApi) List(c *gin.Context) {
	c.JSON(http.StatusOK, commodityLogic.CategoryLogic.List)
}
