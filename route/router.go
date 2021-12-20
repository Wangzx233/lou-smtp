package route

import (
	"github.com/gin-gonic/gin"
	"log"
	"lou-smtp/api"
)

func InitRouter()  {
	router := gin.Default()

	router.GET("/email/send",api.SendEmail)
	err := router.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
