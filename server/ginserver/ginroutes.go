package ginserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func configRoutes(router *gin.Engine, APP_VERSION string) * gin.Engine {
	router.GET("/", func (c *gin.Context){
		c.JSON(http.StatusOK, "HELLO")
	})
	
	return router
}