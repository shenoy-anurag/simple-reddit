package main

import (
	"simple-reddit/configs"
	"simple-reddit/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// load environment variables
	configs.LoadEnvVariables()

	// create gin router server
	router := gin.Default()

	router.Use(cors.Default())
	// ping route as a health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// add various routes to the gin server
	users.Routes(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
