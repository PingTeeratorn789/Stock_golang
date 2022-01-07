package main

import (
	"github.com/gin-gonic/gin"
	"stock/api"
)

func main()  {
	router := gin.Default()
	router.Static("/images", "./uploads/images")

	api.Setup(router)
	gin.SetMode(gin.ReleaseMode)
	router.Run(":8081")
}