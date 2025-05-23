package server

import (
	"github.com/gin-gonic/gin"
	"github.com/velosypedno/genesis-weather-api/internal/ioc"
)

func SetupRoutes(c *ioc.HandlerContainer) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/weather", c.WeatherGETHandler)
		api.POST("/subscribe", c.SubscribePOSTHandler)
		api.GET("/confirm/:token", c.ConfirmGETHandler)
		api.GET("/unsubscribe/:token", c.UnsubscribeGETHandler)
	}
	return router
}
