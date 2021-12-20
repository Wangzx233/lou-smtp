package api

import (
	"github.com/gin-gonic/gin"
	"lou-smtp/service"
)

func SendEmail(c *gin.Context)  {
	name, _ := c.GetQuery("name")
	time, _ := c.GetQuery("time")
	reason, _ := c.GetQuery("reason")

	if name==""||time=="  -"||reason=="" {
		c.JSON(202,gin.H{
			"message":"请输入全部参数",
		})
		return
	}

	err := service.SendEmail(name, time, reason)
	if err != nil {
		c.JSON(202,gin.H{
			"message":"发送失败",
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"发送成功",
	})
}