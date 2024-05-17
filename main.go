package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()
    r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong.GET",
		})
	})
	r.POST("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong.POST",
		})
	})
	r.PUT()
	r.DELETE()
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
