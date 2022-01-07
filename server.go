package main

import (
	"fmt"
	"os"
	"stock/api"

	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	router.Static("/images", "./uploads/images")

	api.Setup(router)
	// router.Run(":8081")

	//In case of running on heroku
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("No Port In Heroku")
		router.Run()
	}else {
		fmt.Println("Enviroment Port : "+ port)
		router.Run(":"+port)
	}
}