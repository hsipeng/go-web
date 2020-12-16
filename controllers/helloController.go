package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "lirawx.cn/go-web/models" // models init
)

// HelloController HelloWorld 接口
// @Summary HelloWorld
// @Description 输出接口
// @Tags hello
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Success 200 {object} _ResponsePostList
// @Router /hello [get]
func HelloController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    "Hello world!",
	})
}
