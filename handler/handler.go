package handler

import "github.com/gin-gonic/gin"

func Handle() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	return r
}
