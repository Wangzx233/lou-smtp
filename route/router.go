package route

import (
	"github.com/gin-gonic/gin"
	"log"
	"lou-smtp/api"
	"lou-smtp/middelware"
)

func InitRouter()  {
	router := gin.Default()
	router.Use(middelware.Cors())
	router.GET("/email/send",api.SendEmail)
	err := router.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
