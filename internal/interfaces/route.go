package interfaces

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPServer(
	us *GinService) *gin.Engine {
	router := gin.New()

	rootGrp := router.Group("/gin/api/v1")
	{
		userGrp := rootGrp.Group("/user")
		userGrp.GET("/sayhi", us.SayHi)
	}

	return router
}
