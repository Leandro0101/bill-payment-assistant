package ginserver

import "github.com/gin-gonic/gin"

type ServerConfig struct {
	PORT string
	APP_VERSION string
}

func Run(config ServerConfig){
	router := configRoutes(gin.Default(), config.APP_VERSION)
	router.Run(":" + config.PORT)
}