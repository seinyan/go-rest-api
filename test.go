package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	server.GET("/", func(context *gin.Context) {

		fmt.Println(context)

	})

	server.Run()
}