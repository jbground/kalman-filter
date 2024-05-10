package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jbground/kalman-filter/api"
)

func RoutingByGin() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	kalman := r.Group("/kalman")
	kalman.GET("/index", api.KalmanIndex)
	kalman.GET("/uniform", api.KalmanUniform)
	kalman.POST("/pdf", api.KalmanPdf)
	kalman.POST("/exec", api.KalmanExecute)
	return r

}
