package main

import (
	"fmt"
	"os"
	"stock/api"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main()  {
	router := gin.Default()

	//Set up CORS middleware options
	config := cors.Config{
		Origins 		: "*",
		RequestHeaders	: "Origin, Authorization, Content-Type",
		Methods 		: "GET, POST, PUT, DELETE",
		Credentials 	: true,
		ValidateHeaders : false,
		MaxAge 			: 1 * time.Minute,
	}

	// Apply the middleware to the router (works on groups too)
	router.Use(cors.Middleware(config))

	router.Static("/images", "./uploads/images")

	api.Setup(router)
	// router.Run(":8081")

	//In case of running on heroku
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("Running on Heroku using random PORT")
		router.Run()
	}else {
		fmt.Println("Enviroment Port : "+ port)
		router.Run(":"+port)
	}
}