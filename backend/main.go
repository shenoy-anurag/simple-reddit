package main

import (
	"simple-reddit/configs"
	"simple-reddit/routes"
)

func main() {
	// load environment variables
	configs.LoadEnvVariables()

	router := routes.SetupRouter()

	router.Run() // listen and serve on 0.0.0.0:8080
}
