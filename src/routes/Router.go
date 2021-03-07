package routes

import (
	controllers "../controllers"
	"github.com/gin-gonic/gin"
)

func MountRoutes(router *gin.RouterGroup) {

	encoderRouter := router.Group("/encode")
	{
		encoderRouter.POST("/base64", controllers.EncodeStringToBase)
	}

	decoderRouter := router.Group("/decode")
	{
		decoderRouter.POST("/base64", controllers.DecodeBaseToString)
	}

}
