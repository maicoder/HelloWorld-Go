package router

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func Init() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/line", func(context *gin.Context) {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			legendData := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
			xAxisData := []int{120, 240, rand.Intn(500), rand.Intn(500), 150, 230, 180}
			context.JSON(200, gin.H{
				"ledeng_data": legendData,
				"xAxis_data": xAxisData,
			})
		})
	}
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error": "404, page not exists!",
		})
	})
	r.Run(":8080")
}
